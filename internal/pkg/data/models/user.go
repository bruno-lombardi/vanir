package models

// @Description User account information
// @Description with user id and email
type User struct {
	ID        string `json:"id" example:"u_AksOKxc12a"`
	Email     string `json:"email" example:"bruno.lombardi@email.com"`
	Name      string `json:"name" example:"Bruno Lombardi"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at" example:"1733583441703"`
	UpdatedAt int64  `json:"updated_at" example:"1733583441710"`
}

// @Description Create user params information
// @Description with email, name and password with confirmation
type CreateUserParams struct {
	Email                string `json:"email" validate:"email,required,max=255" example:"bruno.lombardi@email.com"`
	Name                 string `json:"name" validate:"required,max=100,min=2" example:"Bruno Lombardi"`
	Password             string `json:"password" validate:"required,max=64,min=6" example:"123456"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,max=64,min=6,eqcsfield=Password" example:"123456"`
}

type UpdateUserParams struct {
	ID                      string `param:"id" validate:"required"`
	Email                   string `json:"email" validate:"email,required"`
	Name                    string `json:"name" validate:"required,max=100,min=2"`
	CurrentPassword         string `json:"current_password" validate:"required,max=64,min=6"`
	NewPassword             string `json:"new_password" validate:"required,max=64,min=6"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,max=64,min=6,eqcsfield=NewPassword"`
}

type ListUsersQueryParams struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

type PaginatedUsersResponse struct {
	TotalPages int    `json:"total_pages"`
	Count      int64  `json:"count"`
	PerPage    int    `json:"per_page"`
	Page       int    `json:"page"`
	Data       []User `json:"data"`
}
