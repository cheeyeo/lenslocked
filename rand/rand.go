package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes func in rand: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes fun in rand: didn't read enough random bytes")
	}

	return b, nil
}

// String returns a random string usign crypto/rand
// n is number of bytes being used to generate random string
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
