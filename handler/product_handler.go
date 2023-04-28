package handler

import (
	"strconv"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/service"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (p *productHandler) CreateProduct(c *gin.Context) {
	var productPayload dto.CreateProductRequest

	if err := c.ShouldBindJSON(&productPayload); err != nil {
		bindErr := errs.NewBadRequestError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := p.productService.CreateProduct(productPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (p *productHandler) GetAllProducts(c *gin.Context) {
	response, err := p.productService.GetAllProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Products)
}

func (p *productHandler) UpdateProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())

		c.JSON(paramErr.Status(), paramErr)
		return
	}

	var productPayload dto.PutProductRequest

	if err := c.ShouldBindJSON(&productPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, updateErr := p.productService.UpdateProduct(productId, productPayload)

	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
	}

	c.JSON(response.StatusCode, response)
}

func (p *productHandler) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		paramErr := errs.NewBadRequestError(err.Error())

		c.JSON(paramErr.Status(), paramErr)
		return
	}

	response, deleteErr := p.productService.DeleteProduct(productId)

	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(response.StatusCode, response)
}