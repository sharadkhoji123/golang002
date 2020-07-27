package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "5 Little baby jumping on the bed, one fell down and bumped his head"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("I'm out of options", err)
	}
	fmt.Println(string(bs))
}
