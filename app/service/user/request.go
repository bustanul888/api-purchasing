package user

type userRequest struct{
	UserName string		`json:"user_name" validate:"required"`
	Password string		`json:"password" validate:"required"`
	Role string			`json:"role" validate:"required"`
}

type userUpdateRequest struct{
	UserName string		`json:"user_name" validate:"required"`
	Role string			`json:"role" validate:"required"`
	NewPassword *string		`json:"new_password"`
}

type myProfileUpdateRequest struct {
	UserName string		`json:"user_name" validate:"required"`
	OldPassword string		`json:"old_password"`
	NewPassword *string		`json:"new_password"`
}