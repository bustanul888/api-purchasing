package auth

type authRequest struct {
	UserName 			string `json:"user_name" validate:"required"`
	Password 			string `json:"password" validate:"required,min=7"`
}