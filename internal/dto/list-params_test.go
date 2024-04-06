package dto_test

import (
	"github.com/ak-karimzai/web-labs/internal/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListParams_ValidateNoError(t *testing.T) {
	var listParams dto.ListParams = dto.ListParams{
		PageID:   1,
		PageSize: 5,
	}

	listParams = dto.ListParams{
		PageID:   1000,
		PageSize: 20,
	}
	require.Nil(t, listParams.Validate())
}

func TestListParams_ValidatePageIdIncorrect(t *testing.T) {
	var listParams dto.ListParams = dto.ListParams{
		PageID:   0,
		PageSize: 5,
	}
	require.Error(t, listParams.Validate())
}

func TestListParams_ValidatePageSizeIncorrect(t *testing.T) {
	var listParams dto.ListParams = dto.ListParams{
		PageID:   1,
		PageSize: 4,
	}
	require.Error(t, listParams.Validate())

	listParams = dto.ListParams{
		PageID:   1,
		PageSize: 21,
	}
	require.Error(t, listParams.Validate())
}
