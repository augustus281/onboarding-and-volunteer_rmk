package common

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(99999)%len(letters)]
	}
	return string(b)
}

func GenSalt(len int) string {
	if len < 0 {
		len = 50
	}
	return randSequence(len)
}

func GenRefreshTokenID(len int) string {
	if len < 0 {
		len = 50
	}
	return randSequence(len)
}
