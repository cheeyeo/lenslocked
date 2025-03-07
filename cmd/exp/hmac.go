package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secretKeyForHash := "secret-key"

	password := "this is secret password nobody will know"

	h := hmac.New(sha256.New, []byte(secretKeyForHash))
	h.Write([]byte(password))

	// get the resulting hash
	result := h.Sum(nil)

	fmt.Println(hex.EncodeToString(result))
}
