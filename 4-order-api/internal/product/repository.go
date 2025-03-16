package product

import (
	"4-order-api/pkg/db"
	"errors"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(Db *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: Db,
	}
}

func (repo *ProductRepository) Create(product *Product) error {
	result := repo.Database.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("not founded id")
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetByName(name string) error {
	var product Product
	result := repo.Database.First(&product, "name = ?", name)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return result.Error
	}
	if result.RowsAffected != 0 {
		return errors.New("name already existed")
	}
	return nil
}
