package dto

import "encoding/base64"

type EncryptTextResponse struct {
	Encrypted string `json:"text" binding:"required"`
}

func NewEncryptTextResponse(raw []byte) *EncryptTextResponse {
	encoded := base64.StdEncoding.EncodeToString(raw)
	return &EncryptTextResponse{encoded}
}
