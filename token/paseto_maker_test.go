package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yuttana76/simbplebank/util"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	role := util.DepositorRole
	duration := time.Minute

	issueAt := time.Now()
	expiredAt := issueAt.Add(duration)

	// token, err := maker.CreateToken(username, duration)
	token, payload, err := maker.CreateToken(username, role, duration, TokenTypeAccessToken)

	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// payload, err := maker.VerifyToken(token)
	payload, err = maker.VerifyToken(token, TokenTypeAccessToken)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issueAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString((32)))
	require.NoError(t, err)

	username := util.RandomOwner()
	role := util.DepositorRole

	token, payload, err := maker.CreateToken(username, role, -time.Minute, TokenTypeAccessToken)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token, TokenTypeAccessToken)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

//Homework : add more test for invalid token
