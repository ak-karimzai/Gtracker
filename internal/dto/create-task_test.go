package dto_test

import (
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmptyTask(test *testing.T) {
	var createGoal = dto.CreateTask{}

	require.Error(test, createGoal.Validate())
}

func TestEmptyNameTask(t *testing.T) {
	var createGoal = dto.CreateTask{
		Name:        "",
		Description: util.RandomString(60),
	}

	require.Error(t, createGoal.Validate())
}

func TestEmptyDescriptionTask(t *testing.T) {
	var createGoal = dto.CreateTask{
		Name:        util.RandomString(6),
		Description: "",
		Frequency:   dto.Monthly,
	}

	require.Nil(t, createGoal.Validate())
}
