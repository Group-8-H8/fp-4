package transaction_pg

import (
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/repository/transaction_repository"
	"gorm.io/gorm"
)

// Transaction repository
type transactionRepository struct {
	db *gorm.DB
}

// Transaction repository generator function
func NewPGTransactionRepository(db *gorm.DB) transaction_repository.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

// Create transaction
func (t *transactionRepository) CreateTransaction(transaction entity.Transaction) (*entity.Transaction, errs.Errs) {
	createTransaction := entity.Transaction{
		ProductID: transaction.ProductID,
		UserID: transaction.UserID,
		Quantity: transaction.Quantity,
		TotalPrice: transaction.TotalPrice,
	}

	// Create transaction
	result := t.db.Create(&createTransaction)

	if err := result.Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &createTransaction, nil
}

func (t *transactionRepository) GetMyTransactions(userID int) ([]entity.Transaction, errs.Errs) {
	var myTransactions []entity.Transaction
	
	result := t.db.Model(&myTransactions).Preload("Product").Where("user_id = ?", userID).Find(&myTransactions)

	if err := result.Error; err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}

	return myTransactions, nil
}

func (t *transactionRepository) GetUsersTransactions() ([]entity.Transaction, errs.Errs) {
	var usersTransactions []entity.Transaction

	result := t.db.Model(&usersTransactions).Preload("User").Preload("Product").Find(&usersTransactions)
	if err := result.Error; err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}

	return usersTransactions, nil
}