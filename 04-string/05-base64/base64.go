package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away from keyboard. https://golang.org/"

	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))

	fmt.Println(encodeMessage)

	data, err := base64.StdEncoding.DecodeString(encodeMessage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
		fmt.Println(string(data))
	}
}
