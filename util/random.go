package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomOwnerName() string {
	return randomString(8)
}

func RandomMoney() int64 {
	return randomInt(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "INR", "YEN", "CNY", "HKD"}
	return currencies[rand.Intn(len(currencies))]
}

// Returns a  random number between min and max, inclusive.
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Returns a random String of n characters
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
