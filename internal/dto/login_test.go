package dto_test

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogin_Validate(t *testing.T) {
	var login dto.Login = dto.Login{
		Username: util.RandomString(6),
		Password: util.RandomString(10),
	}

	require.Nil(t, login.Validate())
}
