package rest

import (
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(context *gin.Context) {
	var orders []models.Order
	db.Find(&orders)

	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(http.StatusOK, orders)
}

func CreateOrder(context *gin.Context) {
	var order models.Order
	if err := context.ShouldBind(&order); err != nil {
		context.String(http.StatusBadRequest, "bad request %v", err)
	}
	db.Create(&models.Order{Name: order.Name, Desc: order.Desc, Image: order.Image})

	context.JSON(http.StatusOK, order)
}
