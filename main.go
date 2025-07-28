package main

import (
	"calculadora-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
