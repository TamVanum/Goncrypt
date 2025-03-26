package crypto

type Encryptor interface {
	Encrypt(data []byte) ([]byte, error)
}
