package controller

import (
	"net/http"
	"strconv"

	"github.com/andikanugraha11/go-boilerplate/model/api"
	"github.com/andikanugraha11/go-boilerplate/service"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func InitOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) GetById(c *gin.Context) {
	orderId := c.Param("id")
	id, err := strconv.Atoi(orderId)

	var apiResponse api.ApiResponse
	if err != nil {
		apiResponse = api.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		}
	} else {
		order := controller.OrderService.GetById(c, id)
		apiResponse = api.ApiResponse{
			Code:    http.StatusOK,
			Message: "OK",
			Data:    order,
		}
	}

	c.JSON(apiResponse.Code, apiResponse)

}

func (controller *OrderControllerImpl) GetAll(c *gin.Context) {
	orders := controller.OrderService.GetAll(c)
	apiResponse := api.ApiResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    orders,
	}

	c.JSON(apiResponse.Code, apiResponse)
}

func (controller *OrderControllerImpl) Create(c *gin.Context) {
	orderCreateRequest := api.OrderCreateRequest{}

	c.ShouldBindJSON(&orderCreateRequest)

	orderCreateResponse := controller.OrderService.Create(c, orderCreateRequest)

	apiResponse := api.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    orderCreateResponse,
	}

	c.JSON(apiResponse.Code, apiResponse)
}
