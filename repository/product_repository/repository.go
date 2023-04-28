package product_repository

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
)

// Product repository interface
type ProductRepository interface {
	CreateProduct(product entity.Product) (*entity.Product, errs.Errs)
	GetAllProducts() ([]entity.Product, errs.Errs)
	UpdateProduct(productID int, product entity.Product) (*entity.Product, errs.Errs)
	DeleteProduct(productID int) (string, errs.Errs)
	GetProductById(productID int) (*entity.Product, errs.Errs)
	SellProductById(productID int, quantity int) (*entity.Product, errs.Errs)
}