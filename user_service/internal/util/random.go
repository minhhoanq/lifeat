package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdfeghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RamdomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
