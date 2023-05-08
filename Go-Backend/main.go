package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint `gorm:"primaryKey;unique"`
	Name string
	Age  uint
}

func main() {
	
	r := gin.Default()
	fmt.Print(" 🚀 Server is Up and Running \n\n")

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

	r.POST("/", func(context *gin.Context)  {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		fmt.Println(user)
		context.String(http.StatusOK, "Hello %s", user.Name)
	})

	r.Run(":3000")
}
