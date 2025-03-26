package service

import (
	"strconv"

	"github.com/tamVanum/goncrypt/internal/encryption/crypto"
)

type EncryptionService struct {
	encryptor crypto.Encryptor
}

func NewEncryptionService(enc crypto.Encryptor) *EncryptionService {
	return &EncryptionService{encryptor: enc}
}

func (s *EncryptionService) EncryptText(text string) ([]byte, error) {
	return s.encryptor.Encrypt([]byte(text))
}

func (s *EncryptionService) EncryptNumber(num int) ([]byte, error) {
	return s.encryptor.Encrypt([]byte(strconv.Itoa(num)))
}
