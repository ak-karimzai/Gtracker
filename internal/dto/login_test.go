package dto_test

import (
	"github.com/ak-karimzai/web-labs/internal/dto"
	"github.com/ak-karimzai/web-labs/pkg/util"
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
