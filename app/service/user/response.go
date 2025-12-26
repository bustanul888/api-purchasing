package user

import "task-be/app/model"

type userResponse struct{
	model.BaseModel
	UserName 	string	`json:"user_name"`
	Password 	string	`json:"-"`
	Role		string	`json:"role"`
	model.DateTime
}

func (userResponse) TableName() string{
	return "users"
}