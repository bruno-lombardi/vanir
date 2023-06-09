package models

type AuthCredentials struct {
	Email    string `json:"email" validate:"email,required,max=255"`
	Password string `json:"password" validate:"required,max=64,min=6"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}
