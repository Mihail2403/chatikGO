package hash

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "dncdnjcjdncjdncjdncjdn"
)

// generate password hash
func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
