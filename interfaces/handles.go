package interfaces

type Handles interface {
	MigrationCreateTable()
	SubmitFormData()
	CronjobExcelWriteSheet()
	RewriteExcelWriteSheet()
}
