package user

type userRequest struct{
	UserName string		`json:"user_name" binding:"required"`
	Password string		`json:"password" binding:"required"`
	Role string			`json:"role" binding:"required"`
}

type userUpdateRequest struct{
	UserName string		`json:"user_name" binding:"required"`
	Role string			`json:"role" binding:"required"`
}

type userUpdatePassword struct{
	OldPassword string		`json:"old_password" binding:"required"`
	NewPassword string		`json:"new_password" binding:"required"`
}