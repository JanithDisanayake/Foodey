package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;unique"`
	Name string
	Age  uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		fmt.Println(" 🚀 Database is created \n")
	}
	db.AutoMigrate(&User{})

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
			"id":      id,
		})
	})

	r.POST("/", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		fmt.Println(user)
		context.String(http.StatusOK, "Hello %s", user.Name)
	})

	r.Run(":3000")
}
