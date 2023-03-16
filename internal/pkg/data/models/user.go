package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

type CreateUserDTO struct {
	Email                string `json:"email" validate:"email,required,max=255"`
	Name                 string `json:"name" validate:"required,max=100,min=2"`
	Password             string `json:"password" validate:"required,max=64,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,max=64,min=6,eqcsfield=Password"`
}

type UpdateUserDTO struct {
	ID                      string `param:"id" validate:"required"`
	Email                   string `json:"email" validate:"email,required"`
	Name                    string `json:"name" validate:"required,max=100,min=2"`
	CurrentPassword         string `json:"current_password" validate:"required,max=64,min=6"`
	NewPassword             string `json:"new_password" validate:"required,max=64,min=6"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,max=64,min=6,eqcsfield=NewPassword"`
}

type ListUsersDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

type PaginatedUserResultDTO struct {
	TotalPages int    `json:"total_pages"`
	Count      int64  `json:"count"`
	PerPage    int    `json:"per_page"`
	Page       int    `json:"page"`
	Data       []User `json:"data"`
}
