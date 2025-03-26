package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tamVanum/goncrypt/internal/encryption/handler"
	"github.com/tamVanum/goncrypt/internal/encryption/service"
)

func RegisterEncryptionrRoutes(r *gin.Engine) {

	svc := service.EncryptionService{}
	h := handler.NewEncryptionHandler(&svc)
	r.POST("/encrypt", h.Text)
}
