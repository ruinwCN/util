package ustring

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijkmnpqrstuvwxyzABCDEFGHIJKMNPQRSTUVWXYZ"
const numberBytes = "0123456789"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		j := r.Intn(100)
		m := r.Intn(4)
		if m%2 == 0 {
			b[i] = letterBytes[j%len(letterBytes)]
		} else {
			b[i] = numberBytes[j%len(numberBytes)]
		}
	}
	return string(b)
}
