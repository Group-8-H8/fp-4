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

// CreateProduct godoc
// @Summary Create new product
// @Description Parse request body and add new product data in the database
// @Tags product
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateProductRequest true "Request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 201 {object} dto.CreateProductResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /products [post]
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

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products existing in database
// @Tags product
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {array} dto.GetAllProductsResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /products [get]
func (p *productHandler) GetAllProducts(c *gin.Context) {
	response, err := p.productService.GetAllProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Products)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product by id in url param
// @Tags product
// @Accept json
// @Produce json
// @Param productID  path int true "Product ID"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {object} dto.PutProductResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /products/{productID} [patch]
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

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by id in url param
// @Tags product
// @Produce json
// @Param productID path int true "Product ID"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {object} dto.DeleteProductResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /products/{productID} [delete]
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