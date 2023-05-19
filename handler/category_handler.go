package handler

import (
	"strconv"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/service"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	GetAllCategories(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary Create new category
// @Description Parse request body and add new category data in the database
// @Tags category
// @Accept json
// @Produce json
// @Param RequestBody body dto.CategoryRequest true "Request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 201 {object} dto.CreateCategoryResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /categories [post]
func (ch *categoryHandler) CreateCategory(c *gin.Context) {
	var categoryPayload dto.CategoryRequest

	if err := c.ShouldBindJSON(&categoryPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := ch.categoryService.CreateCategory(categoryPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories and products related existing in database
// @Tags category
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {array} dto.GetAllCategoriesResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /categories [get]
func (ch *categoryHandler) GetAllCategories(c *gin.Context) {
	response, err := ch.categoryService.GetAllCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Categories)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update category by id in url param
// @Tags category
// @Accept json
// @Produce json
// @Param categoryID path int true "Category ID"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {object} dto.PatchCategoryResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /categories/{categoryID} [patch]
func (ch *categoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())

		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var categoryPayload dto.PatchCategoryRequest

	if err := c.ShouldBindJSON(&categoryPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())
		
		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := ch.categoryService.UpdateCategory(id, categoryPayload)

	if err != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	c.JSON(response.StatusCode, response)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete category by id in url param
// @Tags category
// @Produce json
// @Param categoryID  path int true "Category ID"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {object} dto.DeleteCategoryResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /categories/{categoryID} [delete]
func (ch *categoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())

		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := ch.categoryService.DeleteCategory(id)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(response.StatusCode, response)
}