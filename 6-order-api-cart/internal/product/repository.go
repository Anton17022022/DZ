package product

import (
	"6-order-api-cart/pkg/db"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	*db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (repo *ProductRepository) Create(name string) (*Product, error) {
	product, err := repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	if product == nil {
		product = &Product{Name: name}
		result := repo.Db.Create(product)
		if result.Error != nil {
			return nil, err
		}
		return product, nil
	}
	return product, errors.New("product existed")
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	// TODO проверка на наличие
	result := repo.Db.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(name string) error {
	// TODO проверка на наличие
	result := repo.Db.Where("name = ?", name).Delete(&Product{})
	return result.Error
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	result := repo.Db.Where("deleted_at IS NULL").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) FindByName(name string) (*Product, error) {
	var product *Product
	result := repo.Db.First(&product, "name = ?", name)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return product, nil
}
