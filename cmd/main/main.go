package main

import (
	"fmt"

	"github.com/itera-io/crypto-service-go/pkg/crypto"
)

func main() {
	// aes
	// end of aes
	plainText := "1"
	//key := []byte("12345678901234567890123456789012")
	client, _ := crypto.NewClient()
	txt, _ := client.Encrypt(plainText)

	// encryptedData := ""
	// key := []byte("12345678901234567890123456789012")
	// txt, _ := crypto.Decrypt(encryptedData)

	fmt.Println(txt)
}
