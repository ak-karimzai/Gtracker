// Package util provides utility functions for common tasks.
package util

import (
	"math/rand"
	"strings"
)

// alphabet represents the characters used to generate random strings.
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n using characters from the alphabet.
// It returns the generated random string.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
