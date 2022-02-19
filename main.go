package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	n := 16
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error creating random []byte: ", err)
	}

	fmt.Println(base64.URLEncoding.EncodeToString(b))
}
