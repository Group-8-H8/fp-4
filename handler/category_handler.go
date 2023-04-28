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
	GetAllCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

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

func (ch *categoryHandler) GetAllCategory(c *gin.Context) {
	response, err := ch.categoryService.GetAllCategory()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response[0].StatusCode, response)
}

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