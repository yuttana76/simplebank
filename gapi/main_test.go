package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/yuttana76/simbplebank/db/sqlc"
	"github.com/yuttana76/simbplebank/token"
	"github.com/yuttana76/simbplebank/util"
	"github.com/yuttana76/simbplebank/worker"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, username string, duration time.Duration) context.Context {

	role := util.DepositorRole
	accessToken, _, err := tokenMaker.CreateToken(username, role, duration, token.TokenTypeAccessToken)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authoriaztionBearer, accessToken)
	md := metadata.MD{
		authoriaztionHeader: []string{bearerToken},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}
