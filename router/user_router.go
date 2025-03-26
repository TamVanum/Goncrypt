package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tamVanum/goncrypt/internal/user/handler"
	"github.com/tamVanum/goncrypt/internal/user/repository"
	"github.com/tamVanum/goncrypt/internal/user/service"
)

func RegisterUserRoutes(r *gin.Engine) {
	userRepo := repository.NewUserGormRepo()
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	group := r.Group("/users")
	{
		group.POST("/", userHandler.Create)
		group.GET("/", userHandler.GetAll)
	}
}
