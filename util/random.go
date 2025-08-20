package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abshiworhjsnfjelwnfvdf"

// seededRand is a local RNG seeded with current time
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt returns a random integer between min and max inclusive
func RandomInt(min, max int64) int64 {
	return min + seededRand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[seededRand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency picks a random currency
func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR"}
	n := len(currencies)
	return currencies[seededRand.Intn(n)]
}
