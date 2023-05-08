package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	fmt.Println(" 🚀 Server is Up and Running \n")

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World !"})
	})

	r.GET("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		fmt.Println(id)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World !",
			"id" : id,
		})
	})

	r.Run(":3000")
}
