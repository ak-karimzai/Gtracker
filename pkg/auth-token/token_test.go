package auth_token

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	secretKey := util.RandomString(32)
	maker, err := NewJWTToken(secretKey, time.Minute)
	require.NoError(t, err)

	username := util.RandomString(10)
	token, err := maker.CreateToken(10, username)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.Username, username)
	require.True(t, time.Now().Before(payload.ExpiredAt))
}

func TestJWTMakerWrongToken(t *testing.T) {
	secretKey := util.RandomString(32)
	maker, err := NewJWTToken(secretKey, time.Minute)
	require.NoError(t, err)

	username := util.RandomString(10)
	token, err := maker.CreateToken(10, username)
	require.NoError(t, err)

	lastByte := byte(token[len(token)-1]) + 1
	randomToken := token[:len(token)-2] + string(lastByte)
	payload, err := maker.VerifyToken(randomToken)
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestJWTMakerExpiredToken(t *testing.T) {
	secretKey := util.RandomString(32)
	maker, err := NewJWTToken(secretKey, -time.Minute)
	require.NoError(t, err)

	username := util.RandomString(12)
	token, err := maker.CreateToken(10, username)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrInvalidToken)
	require.Nil(t, payload)
}
