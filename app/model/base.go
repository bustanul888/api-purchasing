package model

import (
	"math/rand"
	"task-be/app/helper"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"type:varchar(50)" json:"id"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := ulid.New(ulid.Timestamp(time.Now()), rand.New(rand.NewSource(time.Now().UnixNano())))
	b.ID = id.String()
	return
}

type DateTime struct {
	CreatedAt time.Time      `json:"-" gorm:"type:datetime(0)"`
	UpdatedAt time.Time      `json:"-" gorm:"type:datetime(0)"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:datetime(0)"`
}

func (d *DateTime) BeforeSave(tx *gorm.DB) (err error) {
	d.CreatedAt = helper.UtcTime()
	d.UpdatedAt = helper.UtcTime()
	return
}

func (d *DateTime) BeforeUpdate(tx *gorm.DB) (err error) {
	d.UpdatedAt = helper.UtcTime()
	return
}

func (d *DateTime) BeforeDelete(tx *gorm.DB) (err error) {
	d.DeletedAt.Time = helper.UtcTime()
	return
}
