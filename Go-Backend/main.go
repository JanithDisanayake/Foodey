package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	fmt.Println(" 🚀 Server is Up and Running \n")

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World !"})
	})

	r.Run(":3000")
}
