package handles

import (
	"fmt"

	models "sheet.report.daily/src/migration"
)

func (e *Handles) MigrationCreateTable() {
	var err error
	err = e.Db.Migrator().CreateTable(&models.SubmitFormData{})
	err = e.Backup.Migrator().CreateTable(&models.SubmitFormData{})
	_ = fmt.Sprint(err)
}
