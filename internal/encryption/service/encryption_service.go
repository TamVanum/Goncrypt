package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type EncryptionService struct{}

func (s *EncryptionService) EncryptText(text string) ([]byte, error) {
	byteText := []byte(text)
	key := make([]byte, 32)
	rand.Read(key)

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("cipher init: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, fmt.Errorf("gcm init: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("nonce read: %w", err)
	}

	encrypted := gcm.Seal(nonce, nonce, byteText, nil)
	return encrypted, nil
}
