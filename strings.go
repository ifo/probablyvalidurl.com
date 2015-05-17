package main

import (
	"math/rand"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func setupStrings() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomString(length int) string {
	b, l := make([]rune, length), len(alphabet)
	for i := range b {
		b[i] = alphabet[rand.Intn(l)]
	}
	return string(b)
}
