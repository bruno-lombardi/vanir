package models

import "time"

type Favorite struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Reference string    `json:"reference"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateFavoriteParams struct {
	Type      string `json:"type" validate:"required"`
	Reference string `json:"reference" validate:"required"`
	UserID    string
}

type DeleteFavoriteParams struct {
	Reference string
	UserID    string
}

type ListUserFavoritesQueryParams struct {
	UserID string `validate:"required"`
	Page   int    `query:"page" validate:"required,numeric,min=1"`
	Limit  int    `query:"limit" validate:"required,max=100,min=1"`
}

type ListFavoritesResponse struct {
	Paginated
	Data []Favorite `json:"data"`
}
