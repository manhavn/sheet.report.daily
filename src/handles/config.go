package handles

import (
	"net/http"

	"gorm.io/gorm"
)

type Handles struct {
	Db      *gorm.DB
	Backup  *gorm.DB
	Request *http.Request
	Writer  http.ResponseWriter
}
