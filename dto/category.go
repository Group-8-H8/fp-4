package dto

import "time"

type CategoryRequest struct {
	Type string `json:"type" valid:"required"`
}

type CreateCategoryResponse struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	SoldProductAmount int    `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	StatusCode int `json:"-"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"-"`
}

type PatchCategoryRequest struct {
	Type string `json:"type" valid:"required"` 
}

type PatchCategoryResponse struct {
	ID int `json:"id"`
	Type string `json:"type"`
	SoldProductAmount int `json:"sold_product_amount"`
	UpdatedAt time.Time `json:"updated_at"`
	StatusCode int `json:"-"`
}

type GetCategoryResponse struct {
	ID int `json:"id"`
	Type string `json:"type"`
	SoldProductAmount int `json:"sold_product_amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products []ProductResponseWithoutCategoryId `json:"products"`
	StatusCode int `json:"-"`
}

type GetAllCategoriesResponse struct {
	Categories []GetCategoryResponse
	StatusCode int `json:"-"`
}

