package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "5 Little baby jumping on the bed, one fell down and bumped his head"

	encodeStd :=
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
