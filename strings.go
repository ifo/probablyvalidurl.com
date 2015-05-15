package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed = rand.Reader

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
	rng, err := rand.Int(seed, big.NewInt(int64(x)))
	if err != nil {
		log.Fatal("randomInt: ", err)
	}
	return int(rng.Int64())
}
