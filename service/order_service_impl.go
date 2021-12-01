package service

import (
	"context"

	"github.com/andikanugraha11/go-boilerplate/model/api"
	"github.com/andikanugraha11/go-boilerplate/model/entity"
	"github.com/andikanugraha11/go-boilerplate/repository"

	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *gorm.DB
}

func InitOrderService(orderRepository repository.OrderRepository, DB *gorm.DB) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request api.OrderCreateRequest) api.OrderResponse {

	order := entity.Order{
		Name: request.Name,
	}

	order = service.OrderRepository.Create(ctx, service.DB, order)

	return api.OrderResponse{
		Name: order.Name,
		Id:   order.Id,
	}
}

func (service *OrderServiceImpl) GetAll(ctx context.Context) []api.OrderResponse {

	var orderResponses []api.OrderResponse

	orders := service.OrderRepository.GetAll(ctx, service.DB)

	for _, order := range orders {
		orderResponses = append(orderResponses, api.OrderResponse{Name: order.Name,
			Id: order.Id})
	}

	return orderResponses
}

func (service *OrderServiceImpl) GetById(ctx context.Context, orderId int) api.OrderResponse {
	order, err := service.OrderRepository.GetById(ctx, service.DB, orderId)
	if err != nil {
		return api.OrderResponse{}
	} else {
		return api.OrderResponse{
			Name: order.Name,
			Id:   order.Id,
		}
	}
}
