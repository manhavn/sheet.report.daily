package routers

import (
	"net/http"

	clear_data_expires_handles "sheet.report.daily/src/cronjob"

	"github.com/gorilla/mux"
	"sheet.report.daily/src/handles"
)

func Api(r *mux.Router) {
	// http://localhost:8080/cronjob/set-api-start?api=base64encode_full_url&start=true&spec=*_*_*_*_*
	r.HandleFunc("/cronjob/set-api-start", setApiStart)

	r.HandleFunc("/migration-create-table", migrationCreateTable)
	r.HandleFunc("/submit-form-data", submitFormData)

	r.HandleFunc("/cronjob-excel-write-sheet", cronjobExcelWriteSheet)
	r.HandleFunc("/rewrite-excel-write-sheet", rewriteExcelWriteSheet)

	go func() {
		e := handles.Handles{
			Db:     db.WithContext(cbg),
			Backup: backup.WithContext(cbg),
		}
		e.MigrationCreateTable()

		_, _ = job.AddFunc("*/5 * * * *", func() {
			e := handles.Handles{
				Db:     db.WithContext(cbg),
				Backup: backup.WithContext(cbg),
			}
			e.CronjobExcelWriteSheet()
		})
	}()
}

func migrationCreateTable(writer http.ResponseWriter, request *http.Request) {
	e := handles.Handles{
		Db:      db.WithContext(cbg),
		Backup:  backup.WithContext(cbg),
		Request: request,
		Writer:  writer,
	}
	e.MigrationCreateTable()
}

func submitFormData(writer http.ResponseWriter, request *http.Request) {
	e := handles.Handles{
		Db:      db.WithContext(cbg),
		Backup:  backup.WithContext(cbg),
		Request: request,
		Writer:  writer,
	}
	e.SubmitFormData()
}

func setApiStart(writer http.ResponseWriter, request *http.Request) {
	e := clear_data_expires_handles.CronjobClearDataExpires{
		Db:      db.WithContext(cbg),
		Backup:  backup.WithContext(cbg),
		Request: request,
		Writer:  writer,
	}
	e.CronjobSetApiStart(job, &mapEntry)
}

func cronjobExcelWriteSheet(writer http.ResponseWriter, request *http.Request) {
	e := handles.Handles{
		Db:      db.WithContext(cbg),
		Backup:  backup.WithContext(cbg),
		Request: request,
		Writer:  writer,
	}
	e.CronjobExcelWriteSheet()
}

func rewriteExcelWriteSheet(writer http.ResponseWriter, request *http.Request) {
	e := handles.Handles{
		Db:      db.WithContext(cbg),
		Backup:  backup.WithContext(cbg),
		Request: request,
		Writer:  writer,
	}
	e.RewriteExcelWriteSheet()
}
