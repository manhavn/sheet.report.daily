package routers

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	backup   *gorm.DB
	err      error
	cbg      context.Context
	job      = cron.New()
	mapEntry = map[string]cron.EntryID{}
)

func init() {
	cbg = context.Background()
	db, err = gorm.Open(sqlite.Open("sqlite/gorm.db"), &gorm.Config{})
	if err != nil {
		return
	}
	backup, err = gorm.Open(sqlite.Open("sqlite/backup.db"), &gorm.Config{})
	if err != nil {
		return
	}
	job.Start()
	fmt.Println("Starting Service ...")
}
