package strings

import (
	"math/rand"
	"time"
)

const alphabet string = "-_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Setup() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func SafeRandomString(length int) string {
	return randomString(length, true)
}

func UnsafeRandomString(length int) string {
	return randomString(length, false)
}

func randomString(length int, safe bool) string {
	var res string
	alphalen := len(alphabet)
	if safe {
		alphalen -= 1
	}
	for i := 0; i < length; i++ {
		res += randomChar(randomInt(alphalen), safe)
	}
	return res
}

func randomChar(x int, safe bool) string {
	if safe {
		return alphabet[1:][x : x+1]
	} else {
		return alphabet[x : x+1]
	}
}

func randomInt(x int) int {
	return rand.Intn(x)
}
