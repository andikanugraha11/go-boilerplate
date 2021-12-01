package app

import (
	"github.com/andikanugraha11/go-boilerplate/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(orderController controller.OrderController) *gin.Engine {
	r := gin.Default()

	orderRouter := r.Group("/order")
	{
		orderRouter.GET("/", orderController.GetAll)
		orderRouter.GET("/:id", orderController.GetById)
		orderRouter.POST("/", orderController.Create)
	}

	return r
}
