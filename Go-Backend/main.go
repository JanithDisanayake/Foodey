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

		users := getAllUsers(db)

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, users)
	})

	r.GET("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		var user User
		db.First(&user, "id = ?", id) // find product with integer primary key

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, user)
	})

	r.POST("/", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request %v", err)
		}
		db.Create(&User{Name: user.Name, Age: user.Age})
		fmt.Println(user)
		context.JSON(http.StatusOK, user)
	})

	r.PATCH("/", func(context *gin.Context) {
		var user User
		if err := context.ShouldBind(&user); err != nil {
			context.String(http.StatusBadRequest, "bad request: %v", err)
		}
		user = updateUser(db, user)
		context.JSON(http.StatusOK, user)
	})

	r.DELETE("/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		user := getUserById(db, id)
		db.Delete(&user)
		context.String(http.StatusOK, "user deleted")
	})

	r.Run(":3000")
}

func getAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func getUserById(db *gorm.DB, id int) User {
	var users []User
	db.Find(&users)
	for i := range users {
		if users[i].ID == uint(id) {
			return users[i]
		}
	}
	return User{}
}

func updateUser(db *gorm.DB, user User) User {
	db.Save(user)
	return user
}
