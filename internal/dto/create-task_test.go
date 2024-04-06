package dto_test

import (
	"github.com/ak-karimzai/web-labs/internal/dto"
	"github.com/ak-karimzai/web-labs/pkg/util"
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
