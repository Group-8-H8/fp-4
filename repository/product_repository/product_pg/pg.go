package product_pg

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/repository/product_repository"
	"gorm.io/gorm"
)

// Product repository
type productRepository struct {
	db *gorm.DB
}

// Product repository generator function
func NewPGProductRepository(db *gorm.DB) product_repository.ProductRepository {
	return &productRepository {
		db: db,
	}
}

// Create product
func (p *productRepository) CreateProduct(product entity.Product) (*entity.Product, errs.Errs) {
	// Create new product data
	result := p.db.Create(&product)

	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &product, nil
}

// Get all product
func (p *productRepository) GetAllProducts() ([]entity.Product, errs.Errs) {
	var products []entity.Product

	// Get all product data from database
	result := p.db.Find(&products)

	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return products, nil
}

// Update product
func (p *productRepository) UpdateProduct(productID int, product entity.Product) (*entity.Product, errs.Errs) {
	var productUpdate entity.Product
	
	// Get data by id from database
	result := p.db.First(&productUpdate, productID)
	
	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	// Assign update data
	productUpdate.Title = product.Title
	productUpdate.Price = product.Price
	productUpdate.Stock = product.Stock
	productUpdate.CategoryID = product.CategoryID

	// Update data in database
	result = p.db.Save(&productUpdate)

	err = result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &productUpdate, nil
}

// Delete product
func (p *productRepository) DeleteProduct(productID int) (string, errs.Errs) {
	// Delete product
	result := p.db.Delete(&entity.Product{}, productID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Product has been successfully deleted", nil
}

// Get product by id
func (p *productRepository) GetProductById(productID int) (*entity.Product, errs.Errs) {
	var product entity.Product
	
	// Get product by ID
	result := p.db.First(&product, productID)

	err := result.Error

	if err != nil {
		return nil, errs.NewNotFoundError("product not found")
	}

	return &product, nil
}

// Sell product by id
func (p *productRepository) SellProductById(productId int, quantity int) (*entity.Product, errs.Errs) {
	// Get product by ID
	product, err := p.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	// Reduce stock amount by quantity
	product.Stock -= quantity

	// Update product data in database
	updatedProduct, err := p.UpdateProduct(productId, *product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}