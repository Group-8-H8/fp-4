package dto

import "time"

type ProductResponseWithoutCategoryId struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductResponseWithUpdatedAt struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

type CreateProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	StatusCode int `json:"-"`
}

type GetProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetAllProductResponse struct {
	Products []GetProductResponse
	StatusCode int
}

type PutProductRequest struct {
	Title string `json:"title"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	CategoryId int `json:"category_id"`
}

type PutProductResponse struct {
	Product ProductResponseWithUpdatedAt `json:"product"`
	StatusCode int `json:"-"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"-"`
}