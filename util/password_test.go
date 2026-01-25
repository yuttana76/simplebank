package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPasword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasword1)

	err = CheckPassword(password, hashedPasword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPasword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPasword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasword2)
	require.NotEqual(t, hashedPasword1, hashedPasword2)

}
