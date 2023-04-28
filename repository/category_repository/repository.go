package category_repository

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
)

// Category repository interface
type CategoryRepository interface {
	CreateCategory(categoryType string) (*entity.Category, errs.Errs)
	GetAllCategories() ([]entity.Category, []entity.Product, errs.Errs)
	UpdateCategory(categoryID int, category entity.Category) (*entity.Category, errs.Errs)
	DeleteCategory(categoryID int) (string, errs.Errs)
	GetCategoryById(categoryID int) (*entity.Category, errs.Errs)
}