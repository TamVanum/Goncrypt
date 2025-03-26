package main

import (
	"github.com/tamVanum/goncrypt/config"
	"github.com/tamVanum/goncrypt/router"
)

func main() {
	config.InitDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
