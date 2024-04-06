package dto

import (
	"github.com/ak-karimzai/web-labs/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmptyStruct(test *testing.T) {
	var createGoal = CreateGoal{}

	require.Error(test, createGoal.Validate())
}

func TestEmptyName(t *testing.T) {
	var createGoal = CreateGoal{
		Name:        "",
		Description: util.RandomString(60),
	}

	require.Error(t, createGoal.Validate())
}

func TestEmptyDescription(t *testing.T) {
	var createGoal = CreateGoal{
		Name:        util.RandomString(6),
		Description: "",
		StartDate:   "01-01-2023",
		TargetDate:  "01-01-2024",
	}

	require.Nil(t, createGoal.Validate())
}
