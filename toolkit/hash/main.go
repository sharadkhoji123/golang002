package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test1@example.com")
	fmt.Println(c)
}

func getCode(str string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
