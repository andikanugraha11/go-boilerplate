package service

import (
	"context"

	"github.com/andikanugraha11/go-boilerplate/model/api"
)

type OrderService interface {
	GetAll(ctx context.Context) []api.OrderResponse
	GetById(ctx context.Context, id int) api.OrderResponse
	Create(ctx context.Context, request api.OrderCreateRequest) api.OrderResponse
}
