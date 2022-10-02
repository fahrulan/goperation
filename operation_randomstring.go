package goperation

import (
	"math/rand"
)

func RandomString(length int) string {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	builder := make([]rune, length)
	for i := range builder {
		builder[i] = charset[rand.Intn(len(charset))]
	}

	return string(builder)
}
