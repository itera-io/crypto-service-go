package cryptoservice

type CryptoService interface {
	Encrypt(plainText string) (string, error)
	Decrypt(encryptedData string) (string, error)
}
