package main

import (
	"math/rand"
	"time"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func setup() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomString(length int) string {
	var res string
	for i := 0; i < length; i++ {
		res += randomChar(randomInt(len(alphabet)))
	}
	return res
}

func randomChar(x int) string {
	return alphabet[x : x+1]
}

func randomInt(x int) int {
	return rand.Intn(x)
}
