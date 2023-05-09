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
	r := gin.Default()
	fmt.Print(" 🚀 Server is Up and Running \n\n")

	db, err := gorm.Open(sqlite.Open("foodey.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	} else {
		fmt.Printf(" 🎯 Database is created \n\n")
	}
	db.AutoMigrate(&User{})

	r.GET("/users", func(context *gin.Context) {
		var users []User
		db.Find(&users)

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, users)
	})

	r.GET("/users/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		var user User
		db.First(&user, "id = ?", id) // find product with integer primary key

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, user)
	})

	r.POST("/users", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		db.Create(&User{Name: user.Name, Age: user.Age})

		context.JSON(http.StatusOK, user)
	})

	r.PATCH("/users", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request: %v", err)
		}
		db.Save(user)

		context.JSON(http.StatusOK, user)
	})

	r.DELETE("/users/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		var user User
		db.Take(&user, "id = ?", id)
		db.Delete(&user)

		context.String(http.StatusOK, "user deleted")
	})

	r.Run(":3000")
}
