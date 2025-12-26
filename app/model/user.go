package model

type User struct{
	BaseModel
	UserName	string 	`gorm:"type:varchar(50);uniqueIndex:unique_username"`
	Password string `gorm:"type:varchar(255)"`
	Role	string 	`gorm:"type:varchar(50)"`
	DateTime
}

func (User) TableName() string{
	return "users"
}