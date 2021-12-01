package main

import (
	"github.com/andikanugraha11/go-boilerplate/app"
	"github.com/andikanugraha11/go-boilerplate/controller"
	"github.com/andikanugraha11/go-boilerplate/repository"
	"github.com/andikanugraha11/go-boilerplate/service"
)

func main() {
	db := app.InitDB()

	orderRepository := repository.InitCategoryRepository()
	orderService := service.InitOrderService(orderRepository, db)
	orderController := controller.InitOrderController(orderService)

	r := app.InitRouter(orderController)

	r.Run(":8080")
}
