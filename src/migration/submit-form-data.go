package migration

import (
	"time"
)

type SubmitFormData struct {
	Id        *int32     `gorm:"column:id; type:int; primaryKey; autoIncrement:true;" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at; type:datetime(3);"                 json:"created_at"`
	Month     *string    `gorm:"column:month; type:varchar(10); default:'';"          json:"month"`
	DateNow   *string    `gorm:"column:date_now; type:varchar(12); default:'';"       json:"date_now"`
	TimeNow   *string    `gorm:"column:time_now; type:varchar(10); default:'';"       json:"time_now"`

	FullName *string `gorm:"column:full_name; type:varchar(200); default:'';" json:"full_name"`
	Email    *string `gorm:"column:email; type:varchar(255); default:'';"     json:"email"`
	Phone    *string `gorm:"column:phone; type:varchar(20); default:'';"      json:"phone"`
}

func (SubmitFormData) TableName() string {
	return TableNameSubmitFormData
}

type SubmitFormDataJSON struct {
	Id int32 `gorm:"column:id" json:"id"`

	// CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Month   string `gorm:"column:month"    json:"month"`
	DateNow string `gorm:"column:date_now" json:"date_now"`
	TimeNow string `gorm:"column:time_now" json:"time_now"`

	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email"     json:"email"`
	Phone    string `gorm:"column:phone"     json:"phone"`
}
