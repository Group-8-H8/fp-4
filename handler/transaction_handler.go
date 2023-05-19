package handler

import (
	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/service"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

type TransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetMyTransactions(c *gin.Context)
	GetUsersTransactions(c *gin.Context)
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
	}
}

// CreateTransaction godoc
// @Summary Create new transaction
// @Description Parse request body and add new transaction data in the database
// @Tags transaction
// @Accept json
// @Produce json
// @Param RequestBody body dto.CreateTransactionRequest true "Request body json"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 201 {object} dto.CreateTransactionResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /transactions [post]
func (t *transactionHandler) CreateTransaction(c *gin.Context) {
	user := c.MustGet("userData").(*entity.User)

	userID := int(user.ID)
	
	var transactionPayload dto.CreateTransactionRequest

	if err := c.ShouldBindJSON(&transactionPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	response, err := t.transactionService.CreateTransaction(transactionPayload, userID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

// GetMyTransaction godoc
// @Summary Get logged user transactions
// @Description Get transaction by user id in url param 
// @Tags transaction
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {array} dto.GetMyTransactionsResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /my-transactions [get]
func (t *transactionHandler) GetMyTransactions(c *gin.Context) {
	user := c.MustGet("userData").(*entity.User)

	userID := int(user.ID)

	response, err := t.transactionService.GetMyTransaction(userID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Transactions)
}

// GetUsersTransactions godoc
// @Summary Get all users transactions
// @Description Get transactions by all users 
// @Tags transaction
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
// @Success 200 {array} dto.GetUsersTransactionsResponse
// @Failure 404 {object} errs.Errs
// @Failure 500 {object} errs.Errs
// @Failure 422 {object} errs.Errs
// @Router /user-transactions [get]
func (t *transactionHandler) GetUsersTransactions(c *gin.Context) {
	response, err := t.transactionService.GetUsersTransactions()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Transactions)
}