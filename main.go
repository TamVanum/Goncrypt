package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tamvanum/goncrypt/internal/routes"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.Router(r)
	r.Run(":8080")
}
