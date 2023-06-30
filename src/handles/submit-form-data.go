package handles

import (
	"time"

	"gorm.io/gorm"

	"sheet.report.daily/src/migration"
)

func (e *Handles) SubmitFormData() {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	timeByLoc := time.Now().In(loc)
	month := timeByLoc.Format("2006-01")
	dateNow := timeByLoc.Format("2006-01-02")
	timeNow := timeByLoc.Format("15:04:05")

	fullName := e.Request.FormValue("full_name")
	email := e.Request.FormValue("email")
	phone := e.Request.FormValue("phone")

	go insertData(e.Db, timeByLoc, month, dateNow, timeNow, fullName, email, phone)
	go insertData(e.Backup, timeByLoc, month, dateNow, timeNow, fullName, email, phone)
}

func insertData(
	db *gorm.DB,
	timeByLoc time.Time,
	month string,
	dateNow string,
	timeNow string,
	fullName string,
	email string,
	phone string,
) *gorm.DB {
	return db.Model(&migration.SubmitFormData{}).Create(&migration.SubmitFormData{
		CreatedAt: &timeByLoc,
		Month:     &month,
		DateNow:   &dateNow,
		TimeNow:   &timeNow,
		FullName:  &fullName,
		Email:     &email,
		Phone:     &phone,
	})
}
