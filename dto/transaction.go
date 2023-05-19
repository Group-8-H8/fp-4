package dto

type CreateTransactionRequest struct {
	ProductID int `json:"product_id" valid:"required"`
	Quantity  int `json:"quantity" valid:"required"`
}

type CreateTransactionResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
	StatusCode      int             `json:"-"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type GetMyTransactionsResponse struct {
	Transactions []GetTransactionResponse
	StatusCode   int `json:"-"`
}

type GetTransactionResponse struct {
	ID         int                          `json:"id"`
	ProductID  int                          `json:"product_id"`
	UserID     int                          `json:"user_id"`
	Quantity   int                          `json:"quantity"`
	TotalPrice int                          `json:"total_price"`
	Product    ProductResponseWithUpdatedAt `json:"product"`
}

type GetTransactionResponseWithUser struct {
	ID         int                          `json:"id"`
	ProductID  int                          `json:"product_id"`
	UserID     int                          `json:"user_id"`
	Quantity   int                          `json:"quantity"`
	TotalPrice int                          `json:"total_price"`
	Product    ProductResponseWithUpdatedAt `json:"product"`
	User       UserResponse                 `json:"user"`
}

type GetUsersTransactionsResponse struct {
	Transactions []GetTransactionResponseWithUser
	StatusCode   int `json:"-"`
}
