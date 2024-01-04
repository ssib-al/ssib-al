package randstring

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func Generate(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
