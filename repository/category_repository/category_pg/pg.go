package category_pg

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/repository/category_repository"
	"gorm.io/gorm"
)

// Category repository
type categoryRepository struct {
	db *gorm.DB
}

// Category repository generator function
func NewPGCategoryRepository(db *gorm.DB) category_repository.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

// Create category function
func (c *categoryRepository) CreateCategory(categoryType string) (*entity.Category, errs.Errs) {
	category := entity.Category{
		Type: categoryType,
		SoldProductAmount: 0,
	}

	// Create new category
	result := c.db.Create(&category)

	err := result.Error

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

// Get category by id
func (c *categoryRepository) GetCategoryById(categoryID int) (*entity.Category, errs.Errs) {
	var category entity.Category

	// Get category
	result := c.db.First(&category, categoryID)

	if err := result.Error; err != nil {
		return nil, errs.NewNotFoundError("category not found")
	}

	return &category, nil
}

// Get all category
func (c *categoryRepository) GetAllCategories() ([]entity.Category, []entity.Product, errs.Errs) {
	var categories []entity.Category
	var products []entity.Product

	// Get all category data from database
	result := c.db.Find(&categories)

	err := result.Error

	if err != nil {
		return nil, nil, errs.NewNotFoundError(err.Error())
	}

	var ids []int

	// Collect all category id's
	for _, ctg := range categories {
		ids = append(ids, int(ctg.ID))
	}

	// Get all product with category_id in the id's list
	result = c.db.Where("category_id IN ?", ids).Find(&products)

	err = result.Error

	if err != nil {
		return nil, nil, errs.NewNotFoundError(err.Error())
	}

	return categories, products, nil
}

// Update category
func (c *categoryRepository) UpdateCategory(categoryID int, category entity.Category) (*entity.Category, errs.Errs) {
	var updatedCategory entity.Category

	// Get category by id from database
	result := c.db.First(&updatedCategory, categoryID)

	err := result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	updatedCategory.Type = category.Type
	updatedCategory.SoldProductAmount = category.SoldProductAmount

	// Update data in database
	result = c.db.Save(&updatedCategory)

	err = result.Error

	if err != nil {
		return nil, errs.NewBadRequestError(err.Error())
	}

	return &category, nil
}

// Delete category
func (c *categoryRepository) DeleteCategory(categoryID int) (string, errs.Errs) {
	// Delete data in database
	result := c.db.Delete(&entity.Category{}, categoryID)

	err := result.Error

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return "Category has been successfully deleted", nil
}