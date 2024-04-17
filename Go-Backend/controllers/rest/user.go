package rest

import (
	"go-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = models.New()

func GetAllUsers(context *gin.Context) {
	var users []models.User
	db.Find(&users)

	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(http.StatusOK, users)
}

func GetUserById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user models.User
	db.First(&user, "id = ?", id) // find product with integer primary key

	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "bad request %v", err)
	}
	db.Create(&models.User{Name: user.Name, Age: user.Age})

	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "bad request: %v", err)
	}
	db.Save(user)

	context.JSON(http.StatusOK, user)
}

func RemoveUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user models.User
	db.Take(&user, "id = ?", id)
	db.Delete(&user)

	context.String(http.StatusOK, "user deleted")
}
