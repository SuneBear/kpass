package crypto

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

	password := make([]byte, 0, len(buf))
	// shuffle
	for len(buf) > 1 {
		i := rand.Intn(len(buf))
		switch i {
		case 0:
			password = append(password, buf[0])
			buf = buf[1:]
		case len(buf) - 1:
			password = append(password, buf[i])
			buf = buf[0:i]
		default:
			password = append(password, buf[i])
			buf = append(buf[0:i], buf[i+1:]...)
		}
	}
	return string(append(password, buf[0]))
}
