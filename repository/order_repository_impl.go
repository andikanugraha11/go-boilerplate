package repository

import (
	"context"

	"github.com/andikanugraha11/go-boilerplate/model/entity"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

func InitCategoryRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Create(ctx context.Context, db *gorm.DB, order entity.Order) entity.Order {
	db.Save(order)
	return order
}

func (repository *OrderRepositoryImpl) GetById(ctx context.Context, db *gorm.DB, orderId int) (entity.Order, error) {
	var order entity.Order
	err := db.First(&order, orderId).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (repository *OrderRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) []entity.Order {
	var orders []entity.Order
	db.First(&orders)
	return orders
}
