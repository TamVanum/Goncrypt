package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tamVanum/goncrypt/internal/encryption/crypto"
	"github.com/tamVanum/goncrypt/internal/encryption/handler"
	"github.com/tamVanum/goncrypt/internal/encryption/service"
)

func RegisterEncryptionrRoutes(r *gin.Engine) {

	encryptor := crypto.NewAesGcmEncryptor()
	svc := service.NewEncryptionService(encryptor)
	h := handler.NewEncryptionHandler(svc)
	r.POST("/encrypt", h.Text)
}
