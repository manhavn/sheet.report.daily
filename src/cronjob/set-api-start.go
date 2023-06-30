package clear_data_expires_handles

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/robfig/cron/v3"
)

func (e *CronjobClearDataExpires) CronjobSetApiStart(
	job *cron.Cron,
	mapEntry *map[string]cron.EntryID,
) {
	writer := e.Writer
	request := e.Request
	defer func() {
		sbJob, _ := json.Marshal(mapEntry)
		_, _ = writer.Write(sbJob)
	}()
	list := *mapEntry
	api := strings.TrimSpace(request.FormValue("api"))
	if request.FormValue("start") == "true" {
		if list[api] == 0 {
			spec := "* * * * *"
			customSpec := request.FormValue("spec")
			if customSpec != "" {
				spec = strings.Replace(customSpec, "_", " ", -1)
			}
			entryId, err := job.AddFunc(spec, func() {
				decodedValue, err := base64.StdEncoding.DecodeString(api)
				if err != nil {
					fmt.Println(err)
					return
				}
				go http.Get(string(decodedValue))
			})
			if err == nil {
				list[api] = entryId
			}
		}
	} else {
		entryId := list[api]
		delete(list, api)
		if entryId > 0 {
			job.Remove(entryId)
		}
	}
}
