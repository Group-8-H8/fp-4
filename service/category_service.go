package service

import (
	"net/http"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/pkg/helpers"
	"github.com/fydhfzh/fp-4/repository/category_repository"
)

// Category service struct
type categoryService struct {
	categoryRepo category_repository.CategoryRepository
}

// Category service interface
type CategoryService interface {
	CreateCategory(categoryPayload dto.CategoryRequest) (*dto.CreateCategoryResponse, errs.Errs)
	GetAllCategories() (*dto.GetAllCategoriesResponse, errs.Errs)
	UpdateCategory(categoryId int, categoryPayload dto.PatchCategoryRequest) (*dto.PatchCategoryResponse, errs.Errs)
	DeleteCategory(categoryId int) (*dto.DeleteCategoryResponse, errs.Errs)
}

// Category service generator function
func NewCategoryService(categoryRepo category_repository.CategoryRepository) (CategoryService) {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

// Create category service
func (c *categoryService) CreateCategory(categoryPayload dto.CategoryRequest) (*dto.CreateCategoryResponse, errs.Errs) {
	// Validate category payload struct
	err := helpers.ValidateStruct(categoryPayload)

	if err != nil {
		return nil, err
	}

	// Create category
	category, err := c.categoryRepo.CreateCategory(categoryPayload.Type)

	if err != nil {
		return nil, err
	}

	// Create category response
	response := dto.CreateCategoryResponse{
		ID: int(category.ID),
		Type: category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt: category.CreatedAt,
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

// Get all category service
func (c *categoryService) GetAllCategories() (*dto.GetAllCategoriesResponse, errs.Errs) {
	// Get all categories
	categories, products, err := c.categoryRepo.GetAllCategories()

	if err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}

	// Create get category response
	var categoriesResponse []dto.GetCategoryResponse

	for _, category := range categories {
		categoryResponse := dto.GetCategoryResponse{
			ID: int(category.ID),
			Type: category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			StatusCode: http.StatusOK,
		}

		var productsResponse []dto.ProductResponseWithoutCategoryId

		for _, product := range products {
			if product.CategoryID == int(category.ID) {
				productResponse := dto.ProductResponseWithoutCategoryId{
					ID: int(product.ID),
					Title: product.Title,
					Price: product.Price,
					Stock: product.Stock,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
				}

				productsResponse = append(productsResponse, productResponse)
			}
		}

		categoryResponse.Products = productsResponse

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	response := dto.GetAllCategoriesResponse{
		Categories: categoriesResponse,
		StatusCode: http.StatusOK,
	}

	return &response, nil
}

// Update category service
func (c *categoryService) UpdateCategory(categoryId int, categoryPayload dto.PatchCategoryRequest) (*dto.PatchCategoryResponse, errs.Errs) {
	// Validate category payload structs
	err := helpers.ValidateStruct(categoryPayload)

	if err != nil {
		return nil, err
	}

	updateCategory := entity.Category{
		Type: categoryPayload.Type,
	}
	
	// Update category
	category, err := c.categoryRepo.UpdateCategory(categoryId, updateCategory)

	if err != nil {
		return nil, err
	}

	// Create update category response
	response := dto.PatchCategoryResponse{
		ID: int(category.ID),
		Type: category.Type,
		SoldProductAmount: category.SoldProductAmount,
		UpdatedAt: category.UpdatedAt,
		StatusCode: http.StatusOK,
	}

	return &response, nil
}

// Delete category service
func (c *categoryService) DeleteCategory(id int) (*dto.DeleteCategoryResponse, errs.Errs) {
	// Delete category
	message, err := c.categoryRepo.DeleteCategory(id)

	if err != nil {
		return nil, err
	}

	// Create delete category response
	response := dto.DeleteCategoryResponse{
		Message: message,
		StatusCode: http.StatusOK,
	}

	return &response, err
}