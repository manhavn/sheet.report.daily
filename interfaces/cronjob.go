package interfaces

import "github.com/robfig/cron/v3"

type CronjobClearDataExpires interface {
	CronjobSetApiStart(*cron.Cron, *map[string]cron.EntryID)
}
