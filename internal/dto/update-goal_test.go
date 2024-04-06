package dto_test

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateGoal_NothingToUpdate(t *testing.T) {
	var updateGoal dto.UpdateGoal
	require.EqualError(t, updateGoal.Validate(), dto.ErrEmptyUpdate.Error())
}
