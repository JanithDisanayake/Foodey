package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	// Create a new Gin router
	router := gin.Default()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	router.Run()
}
