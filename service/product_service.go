package service

import (
	"net/http"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/pkg/helpers"
	"github.com/fydhfzh/fp-4/repository/product_repository"
)

// Product service struct
type productService struct {
	productRepo product_repository.ProductRepository
}

// Product service interface
type ProductService interface {
	CreateProduct(productPayload dto.CreateProductRequest) (*dto.CreateProductResponse, errs.Errs) 
	GetAllProducts() (*dto.GetAllProductsResponse, errs.Errs) 
	UpdateProduct(productId int, productPayload dto.PutProductRequest) (*dto.PutProductResponse, errs.Errs)
	DeleteProduct(productId int) (*dto.DeleteProductResponse, errs.Errs)
}

// Product service generator function
func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

// Create product service
func (p *productService) CreateProduct(productPayload dto.CreateProductRequest) (*dto.CreateProductResponse, errs.Errs) {
	// Validate product payload struct
	err := helpers.ValidateStruct(productPayload)

	if err != nil {
		return nil, err
	}
	
	// Create product entity
	product := entity.Product{
		Title: productPayload.Title,
		Price: productPayload.Price,
		Stock: productPayload.Stock,
		CategoryID: productPayload.CategoryID,
	}

	productCreated, err := p.productRepo.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	// Create product response
	response := dto.CreateProductResponse{
		ID: int(productCreated.ID),
		Title: productCreated.Title,
		Price: productCreated.Price,
		Stock: productCreated.Stock,
		CategoryID: productCreated.CategoryID,
		CreatedAt: productCreated.CreatedAt,
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

// Get all products service
func (p *productService) GetAllProducts() (*dto.GetAllProductsResponse, errs.Errs) {
	// Get all products
	products, err := p.productRepo.GetAllProducts()

	if err != nil {
		return nil, err
	}

	// Create get all products response
	var productsResponse dto.GetAllProductsResponse
	
	for _, product := range products {
		productResponse := dto.GetProductResponse{
			ID: int(product.ID),
			Title: product.Title,
			Price: product.Price,
			Stock: product.Stock,
			CategoryId: product.CategoryID,
			CreatedAt: product.CreatedAt,
		}

		productsResponse.Products = append(productsResponse.Products, productResponse)
	}

	productsResponse.StatusCode = http.StatusOK

	return &productsResponse, nil
}

// Update product service
func (p *productService) UpdateProduct(productId int, productPayload dto.PutProductRequest) (*dto.PutProductResponse, errs.Errs) {
	// Validate product payload struct
	err := helpers.ValidateStruct(productPayload)

	if err != nil {
		return nil, err
	}
	
	// Create product entity
	product := entity.Product {
		Title: productPayload.Title,
		Price: productPayload.Price,
		Stock: productPayload.Stock,
		CategoryID: productPayload.CategoryId,
	}

	// Update product
	productUpdated, err := p.productRepo.UpdateProduct(productId, product)

	if err != nil {
		return nil, err
	}

	// Create update product response
	response := dto.PutProductResponse{
		Product: dto.ProductResponseWithUpdatedAt{
			ID: int(productUpdated.ID),
			Title: productUpdated.Title,
			Price: productUpdated.Price,
			Stock: productUpdated.Stock,
			CategoryID: productUpdated.CategoryID,
			CreatedAt: productUpdated.CreatedAt,
			UpdatedAt: productUpdated.UpdatedAt,
		},
		StatusCode: http.StatusOK,
	}

	return &response, nil
}

// Delete product service
func (p *productService) DeleteProduct(productId int) (*dto.DeleteProductResponse, errs.Errs) {
	// Delete product
	message, err := p.productRepo.DeleteProduct(productId)

	if err != nil {
		return nil, err
	}

	// Create delete product response
	response := dto.DeleteProductResponse{
		Message: message,
		StatusCode: http.StatusOK,
	}

	return &response, nil
}