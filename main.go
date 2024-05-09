package main

import (
	"log"
	"os"

	"examplegpt.com/router"
	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	router.Router(serv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3033"
	}
	if err := serv.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
