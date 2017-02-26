package util

import (
	"math/rand"
	"time"
)

const (
	numericTables = "0123456789"
	specialTables = "!#$%&()*+,-_./:;=?@[]^{}~|"
	simpleTables  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RandPass ...
func RandPass(length int, count ...int) string {
	if length < 4 {
		length = 4
	} else if length > 64 {
		length = 64
	}
	numCount := 0
	specCount := 0
	if len(count) > 0 && count[0] > 0 && count[0] <= 10 {
		numCount = count[0]
	}
	if len(count) > 1 && count[1] > 0 && count[1] <= 10 {
		specCount = count[1]
	}

	buf := make([]byte, 0, length)
	rand.Seed(time.Now().UnixNano())
	for len(buf) < specCount {
		buf = append(buf, specialTables[rand.Intn(len(specialTables))])
	}
	for len(buf) < specCount+numCount {
		buf = append(buf, numericTables[rand.Intn(len(numericTables))])
	}
	for len(buf) < length {
		buf = append(buf, simpleTables[rand.Intn(len(simpleTables))])
	}

	password := make([]byte, length)
	// shuffle
	for i, j := range rand.Perm(length) {
		password[i] = buf[j]
	}
	return string(password)
}
