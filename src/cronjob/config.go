package clear_data_expires_handles

import (
	"net/http"

	"gorm.io/gorm"
)

type CronjobClearDataExpires struct {
	Db      *gorm.DB
	Backup  *gorm.DB
	Request *http.Request
	Writer  http.ResponseWriter
}
