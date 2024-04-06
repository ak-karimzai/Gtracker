package dto

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDDMMYYYY_Validate(t *testing.T) {
	date := DDMMYYYY("29022000")
	isValid := date.Validate()
	require.False(t, isValid)

	date = DDMMYYYY("29--33-2000")
	isValid = date.Validate()
	require.False(t, isValid)
}

func TestCorrect(t *testing.T) {
	stringDate := "01-01-2023"
	date := DDMMYYYY(stringDate)
	require.True(t, date.Validate())

	grigDate := date.ToStdDate()
	require.Equal(t, grigDate.Year(), 2023)
	require.Equal(t, grigDate.Month(), time.Month(1))
	require.Equal(t, grigDate.Day(), 01)
}

func TestCorrectDate(t *testing.T) {
	stringDate := "03-03-2024"
	date := DDMMYYYY(stringDate)
	require.True(t, date.Validate())

	grigDate := date.ToStdDate()
	require.Equal(t, grigDate.Year(), 2024)
	require.Equal(t, grigDate.Month(), time.Month(3))
	require.Equal(t, grigDate.Day(), 03)
}
