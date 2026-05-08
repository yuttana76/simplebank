package gapi

import (
	"context"
	"database/sql"

	db "github.com/yuttana76/simbplebank/db/sqlc"
	"github.com/yuttana76/simbplebank/pb"
	"github.com/yuttana76/simbplebank/token"
	"github.com/yuttana76/simbplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	//Step 1: Get user from database
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	//Step 2: Check password
	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	role := util.DepositorRole
	//Step 3: Create access token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		role,
		server.config.AccessTokenDuration,
		token.TokenTypeAccessToken,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	//Step 4: Create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		role,
		server.config.RefreshTokenDuration,
		token.TokenTypeRefreshToken,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	//Step 5: Create session in database
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil
}
