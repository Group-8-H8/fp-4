package transaction_repository

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
)

// Transaction repository interface
type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (*entity.Transaction, errs.Errs)
	GetMyTransactions(userId int) ([]entity.Transaction, errs.Errs)
	GetUsersTransactions() ([]entity.Transaction, errs.Errs)
}