package dto_test

import (
	"github.com/ak-karimzai/web-labs/internal/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFrequency_Validate(t *testing.T) {
	var frequency dto.Frequency = dto.Frequency("Daily")
	require.True(t, frequency.Validate())

	frequency = dto.Frequency("Weekly")
	require.True(t, frequency.Validate())

	frequency = dto.Frequency("Monthly")
	require.True(t, frequency.Validate())
}
