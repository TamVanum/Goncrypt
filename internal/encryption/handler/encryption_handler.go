package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamVanum/goncrypt/internal/encryption/dto"
	"github.com/tamVanum/goncrypt/internal/encryption/service"
)

type EncryptionHandler struct {
	svc *service.EncryptionService
}

func NewEncryptionHandler(svc *service.EncryptionService) *EncryptionHandler {
	return &EncryptionHandler{svc}
}

func (h *EncryptionHandler) Text(c *gin.Context) {
	var req dto.EncryptTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text field required"})
		return
	}

	encrypted, err := h.svc.EncryptText(req.Text)
	if err != nil {
		log.Printf("encryption error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"encrypted": dto.NewEncryptTextResponse(encrypted)})
}

func (h *EncryptionHandler) Number(c *gin.Context) {
	var req dto.EncryptNumberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Number Encryption Failed"})
		return
	}

	encrypted, err := h.svc.EncryptNumber(req.Number)

	if err != nil {
		log.Printf("encryption error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption Failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"encrypted": encrypted})
}
