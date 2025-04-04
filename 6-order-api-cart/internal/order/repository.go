package order

import (
	"6-order-api-cart/internal/product"
	"6-order-api-cart/pkg/db"
	"errors"
)

type OrderRepository struct {
	*db.Db
	*product.ProductRepository
}

type OrderRepositoryDeps struct {
	*db.Db
	*product.ProductRepository
}

func NewOrderRepository(deps *OrderRepositoryDeps) *OrderRepository {
	return &OrderRepository{
		Db:                deps.Db,
		ProductRepository: deps.ProductRepository,
	}
}

func (repo *OrderRepository) Create(order *Order) error {
	// TODO проверка на наличие остатков
	product, err := repo.ProductRepository.FindByName(order.ProductName)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not existed")
	}

	result := repo.Db.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *OrderRepository) GetOrderByID(id uint) (*Order, error){
	var order *Order
	result := repo.Db.First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetOrdersByPhoneNumber(usePhoneNumber string)([]Order, error){
	var orders []Order
	result := repo.Db.Where("deleted_at IS NULL").Where("user_phone_number = ?", usePhoneNumber).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
