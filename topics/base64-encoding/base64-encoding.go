package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "test base 64 data"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sDec))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(uDec))
}
