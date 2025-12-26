package model

type BlackListToken struct {
	BaseModel
	Token string `gorm:"type:varchar(255)"`
	DateTime
}

func (BlackListToken) TableName() string {
	return "blacklist_tokens"
}
