package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type AesGcmEncryptor struct {
}

func NewAesGcmEncryptor() *AesGcmEncryptor {
	return &AesGcmEncryptor{}
}

func (e *AesGcmEncryptor) Encrypt(data []byte) ([]byte, error) {
	key := make([]byte, 32)
	rand.Read(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("cipher init: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("gcm init: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("nonce read: %w", err)
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}
