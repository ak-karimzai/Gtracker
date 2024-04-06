package dto_test

import (
	"github.com/ak-karimzai/web-labs/internal/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateGoal_NothingToUpdate(t *testing.T) {
	var updateGoal dto.UpdateGoal
	require.EqualError(t, updateGoal.Validate(), dto.ErrEmptyUpdate.Error())
}
