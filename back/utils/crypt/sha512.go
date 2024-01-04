package crypt

import (
	"crypto/sha512"
	"fmt"
)

func SHA512(raw string) (crypted string) {
	crypted = fmt.Sprintf("%x", sha512.Sum512([]byte(raw)))
	return
}
