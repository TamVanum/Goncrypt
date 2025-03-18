package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamvanum/goncrypt/internal/handlers"
)

func Router(r *gin.Engine) {
	r.GET("/", handlers.HealthCheck)
}
