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

func (t *transactionHandler) GetUsersTransactions(c *gin.Context) {
	response, err := t.transactionService.GetUsersTransactions()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response.Transactions)
}