package dto

type EncryptTextRequest struct {
	Text string `json:"text" binding:"required"`
}

type EncryptNumberRequest struct {
	Number int `json:"number" binding:"required"`
}
