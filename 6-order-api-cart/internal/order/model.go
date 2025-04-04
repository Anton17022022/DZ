package order

import (
	"6-order-api-cart/internal/product"
	"6-order-api-cart/internal/user"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductName     string `json:"productname"`
	Product         product.Product `gorm:"foreignKey:ProductName;references:Name"`
	Amount          string    `json:"amount"`
	UserPhoneNumber string    `gorm:"foreignKey:ProductName"`
	User            user.User `gorm:"foreignKey:UserPhoneNumber;references:PhoneNumber"`
}

func NewOrder(product, PhoneNumber, amount string) *Order {
	order := &Order{
		ProductName:     product,
		Amount:          amount,
		UserPhoneNumber: PhoneNumber,
	}
	return order
}
