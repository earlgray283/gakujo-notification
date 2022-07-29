package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// ref: https://deeeet.com/writing/2015/11/10/go-crypto/
type Crypto struct {
	block cipher.Block
}

func NewCrypto(key []byte) (*Crypto, error) {
	if len(key) == 0 {
		return nil, errors.New("the environmental value GAKUJO_NOTIFICATION_ENCRYPT_KEY must be set")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &Crypto{block}, nil
}

func (c *Crypto) Encrypt(text []byte) (string, error) {
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCTR(c.block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], text)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (c *Crypto) Decrypt(base64CipherText string) (string, error) {
	cipherText := make([]byte, base64.StdEncoding.DecodedLen(len(base64CipherText)))
	if _, err := base64.StdEncoding.Decode(cipherText, []byte(base64CipherText)); err != nil {
		return "", err
	}
	decryptedText := make([]byte, len(cipherText[aes.BlockSize:]))
	stream := cipher.NewCTR(c.block, cipherText[:aes.BlockSize])
	stream.XORKeyStream(decryptedText, cipherText[aes.BlockSize:])
	return string(decryptedText), nil
}
