package service

import (
	"net/http"

	"github.com/fydhfzh/fp-4/dto"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
	"github.com/fydhfzh/fp-4/repository/category_repository"
	"github.com/fydhfzh/fp-4/repository/product_repository"
	"github.com/fydhfzh/fp-4/repository/transaction_repository"
	"github.com/fydhfzh/fp-4/repository/user_repository"
)

// Transaction service struct
type transactionService struct {
	transactionRepo transaction_repository.TransactionRepository
	productRepo product_repository.ProductRepository
	userRepo user_repository.UserRepository
	categoryRepo category_repository.CategoryRepository
}

// Transaction service interface
type TransactionService interface {
	CreateTransaction(transactionPayload dto.CreateTransactionRequest, userID int) (*dto.CreateTransactionResponse, errs.Errs) 
	GetMyTransaction(userID int) (*dto.GetMyTransactionsResponse, errs.Errs)
	GetUsersTransactions() (*dto.GetUsersTransactionsResponse, errs.Errs)
}

// Transaction service generator function
func NewTransactionService(
	transactionRepo transaction_repository.TransactionRepository,
	productRepo product_repository.ProductRepository,
	userRepo user_repository.UserRepository,
	categoryRepo category_repository.CategoryRepository,
	) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		productRepo: productRepo,
		userRepo: userRepo,
		categoryRepo: categoryRepo,
	}
}

// Create transaction service
func (t *transactionService) CreateTransaction(transactionPayload dto.CreateTransactionRequest, userID int) (*dto.CreateTransactionResponse, errs.Errs) {
	// Create transaction entity
	transaction := entity.Transaction{
		ProductID: transactionPayload.ProductID,
		UserID: userID,
		Quantity: transactionPayload.Quantity,
	}

	// Get product data by id if exists
	product, err := t.productRepo.GetProductById(transaction.ProductID)

	if err != nil {
		return nil, err
	}

	// Get user data by id
	user, err := t.userRepo.GetUserById(transaction.UserID)

	if err != nil {
		return nil, err
	}

	// Get category data by id
	category, err := t.categoryRepo.GetCategoryById(product.CategoryID)

	if err != nil {
		return nil, err
	}

	// Calculate total price
	transaction.TotalPrice = product.Price * transaction.Quantity

	// Check if total price is less than user balance
	if transaction.TotalPrice > user.Balance {
		return nil, errs.NewBadRequestError("your balance is not enough")
	}

	// Check if quantity is less than product stock available
	if transaction.Quantity > product.Stock {
		return nil, errs.NewBadRequestError("product stock is less than demand")
	}
	
	// Sell product
	product, err = t.productRepo.SellProductById(transaction.ProductID, transaction.Quantity)

	if err != nil {
		return nil, err
	}

	// Pay order
	err = t.userRepo.PayOrder(transaction.TotalPrice, transaction.UserID)

	if err != nil {
		return nil, err
	}

	updateCategory := entity.Category {
		Type: category.Type,
		SoldProductAmount: category.SoldProductAmount + transaction.Quantity,
	}
	
	// Update category sold_product_amount field
	_, err = t.categoryRepo.UpdateCategory(product.CategoryID, updateCategory)

	if err != nil {
		return nil, err
	}

	// Create transaction
	transactionCreated, err := t.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}

	// Create transaction response
	response := dto.CreateTransactionResponse{
		Message: "You have successfully purchased the product",
		TransactionBill: dto.TransactionBill{
			TotalPrice: transactionCreated.TotalPrice,
			Quantity: transactionCreated.Quantity,
			ProductTitle: product.Title,
		},
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

func (t *transactionService) GetMyTransaction(userID int) (*dto.GetMyTransactionsResponse, errs.Errs) {
	myTransactions, err := t.transactionRepo.GetMyTransactions(userID)

	if err != nil {
		return nil, err
	}

	var transactions []dto.GetTransactionResponse

	for _, tx := range myTransactions {
		product := dto.ProductResponseWithUpdatedAt{
			ID: tx.ProductID,
			Title: tx.Product.Title,
			Price: tx.Product.Price,
			Stock: tx.Product.Stock,
			CategoryID: tx.Product.CategoryID,
			CreatedAt: tx.Product.CreatedAt,
			UpdatedAt: tx.Product.UpdatedAt,
		}

		transaction := dto.GetTransactionResponse{
			ID: int(tx.ID),
			ProductID: tx.ProductID,
			UserID: tx.UserID,
			Quantity: tx.Quantity,
			TotalPrice: tx.TotalPrice,
			Product: product,
		}
		
		transactions = append(transactions, transaction)
	}

	response := dto.GetMyTransactionsResponse{
		Transactions: transactions,
		StatusCode: http.StatusOK,
	}

	return &response, nil
}

func (t *transactionService) GetUsersTransactions() (*dto.GetUsersTransactionsResponse, errs.Errs) {
	usersTransactions, err := t.transactionRepo.GetUsersTransactions()

	if err != nil {
		return nil, err
	}

	var transactionsResponse dto.GetUsersTransactionsResponse

	for _, tx := range usersTransactions {
		product := dto.ProductResponseWithUpdatedAt{
			ID: int(tx.Product.ID),
			Title: tx.Product.Title,
			Price: tx.Product.Price,
			Stock: tx.Product.Stock,
			CategoryID: tx.Product.CategoryID,
			CreatedAt: tx.Product.CreatedAt,
			UpdatedAt: tx.Product.UpdatedAt,
		}

		user := dto.UserResponse{
			ID: int(tx.User.ID),
			Email: tx.User.Email,
			Fullname: tx.User.FullName,
			Balance: tx.User.Balance,
			CreatedAt: tx.User.CreatedAt,
			UpdatedAt: tx.User.UpdatedAt,
		}

		transaction := dto.GetTransactionResponseWithUser{
			ID: int(tx.ID),
			ProductID: tx.ProductID,
			UserID: tx.UserID,
			Quantity: tx.Quantity,
			TotalPrice: tx.TotalPrice,
			Product: product,
			User: user,
		}

		transactionsResponse.Transactions = append(transactionsResponse.Transactions, transaction)
	}

	transactionsResponse.StatusCode = http.StatusOK

	return &transactionsResponse, nil
}