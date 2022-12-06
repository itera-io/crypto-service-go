package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/itera-io/crypto-service-go/pkg/cryptoservice"
)

type cryptoClient struct {
	Key []byte
}

const KeySize = 256
const CryptoServiceKeyEnvVar = "CRYPTO_SERVICE_KEY"

func NewClient() (cryptoservice.CryptoService, error) {
	key := os.Getenv(CryptoServiceKeyEnvVar)
	if key == "" {
		return nil, fmt.Errorf(
			`please specify crypto service key %s`,
			CryptoServiceKeyEnvVar)
	}
	return &cryptoClient{Key: []byte(key)}, nil
}

func (c *cryptoClient) Encrypt(plainText string) (string, error) {
	// Generate random 16 bytes for symmetric algorithm vector
	iv, _ := GenerateRandomBytes(KeySize / 16)

	if len(iv) != 16 {
		fmt.Printf("IV len must be 16. its: %v\n", len(iv))
	}
	var encrypted string
	toEncodeByte := []byte(plainText)
	toEncodeBytePadded := PKCS5Padding(toEncodeByte, aes.BlockSize)

	block, err := aes.NewCipher(c.Key)
	if err != nil {
		fmt.Printf("CBC Encrypting failed %s", err.Error())
		return "", fmt.Errorf("CBC Encrypting failed")
	}
	ciphertext := make([]byte, len(toEncodeBytePadded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, toEncodeBytePadded)
	encrypted = hex.EncodeToString(iv) + hex.EncodeToString(ciphertext)
	return encrypted, nil
}

func (c *cryptoClient) Decrypt(encryptedData string) (string, error) {

	iv, err := hex.DecodeString(encryptedData[:KeySize/8])
	if err != nil {
		return "", fmt.Errorf("specify valid encrypted data")
	}
	encryptedData = encryptedData[KeySize/8:]

	if len(iv) != 16 {
		fmt.Printf("IV len must be 16. its: %v\n", len(iv))
	}
	txt, _ := hex.DecodeString(encryptedData)
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		fmt.Printf("CBC Decrypting failed. %s", err.Error())
		return "", fmt.Errorf("CBC Decrypting failed")
	}
	ciphertext := make([]byte, len(txt))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, txt)
	ciphertext = PKCS7UnPadding(ciphertext)
	return string(ciphertext), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}
