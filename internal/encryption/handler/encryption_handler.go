package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamVanum/goncrypt/internal/encryption/service"
)

type EncryptionHandler struct {
	svc *service.EncryptionService
}

func NewEncryptionHandler(svc *service.EncryptionService) *EncryptionHandler {
	return &EncryptionHandler{svc}
}

type EncryptRequest struct {
	Text string `json:"text" binding:"required"`
}

func (h *EncryptionHandler) Text(c *gin.Context) {
	var req EncryptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text field required"})
		return
	}

	encrypted, err := h.svc.EncryptText(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"encrypted": encrypted})

}
