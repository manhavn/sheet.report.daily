package handles

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"sheet.report.daily/excel"
	"sheet.report.daily/src/migration"
)

func (e *Handles) CronjobExcelWriteSheet() {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	timeNow := time.Now().In(loc)
	timeYesterday := timeNow.AddDate(0, 0, 0)
	sheetName := timeYesterday.Format("2006-01") // 2023-07
	timeInLastHour := time.Date(
		timeNow.Year(),
		timeNow.Month(),
		timeNow.Day(),
		timeNow.Hour(),
		0,
		0,
		0,
		loc,
	).Add(time.Hour * -1)

	var token []byte
	var credentials []byte
	var sheetId string
	var sheetIdShow string

	func() { // config
		if e.Request != nil {
			customMonth := strings.TrimSpace(e.Request.FormValue("custom_month")) // 2023-06
			if len(customMonth) == 7 {
				sheetName = customMonth
			}
		}

		sheetReader, _ := os.ReadFile("sheet_connect/sheet.json")
		var sheetConfigId struct {
			Id     string `json:"id"`
			IdShow string `json:"id_show"`
		}
		_ = json.Unmarshal(sheetReader, &sheetConfigId)
		sheetId = sheetConfigId.Id
		sheetIdShow = sheetConfigId.IdShow

		var err error
		credentials, err = os.ReadFile("sheet_connect/credentials.json")
		if err == nil {
			tokenJson := "sheet_connect/token.json"
			if tokenReader, err := os.ReadFile(tokenJson); err == nil {
				var gvt excel.GoogleVerificationsToken
				_ = json.Unmarshal(tokenReader, &gvt)
				token = excel.RefreshTokenSheet(&gvt, credentials, tokenJson)
			}
		}
	}()

	go func(credentials []byte, token []byte, sheetId string, sheetIdShow string, sheetName string) {
		var newData [][]interface{}
		page := 0
		limit := 50
		var maxId int32
		for {
			var found []migration.SubmitFormDataJSON
			e.Db.Model(&migration.SubmitFormData{}).
				Where(&migration.SubmitFormData{Month: &sheetName}).
				Where("created_at < ?", timeInLastHour).
				Offset(page * limit).Limit(limit).
				Find(&found)
			page++
			var next bool
			for _, row := range found {
				next = true
				if maxId < row.Id {
					maxId = row.Id
				}
				newData = append(newData, []interface{}{
					"",
					row.DateNow,
					row.TimeNow,
					row.FullName,
					row.Phone,
					row.Email,
				})
			}
			if !next {
				break
			}
			time.Sleep(time.Second / 5)
		}

		if excel.SheetWriteFile(
			credentials,
			token,
			sheetId,
			sheetIdShow,
			sheetName,
			newData,
		) == nil {
			e.Db.Model(&migration.SubmitFormData{}).Where("id <= ?", maxId).Delete(nil)
		}
	}(
		credentials,
		token,
		sheetId,
		sheetIdShow,
		sheetName,
	)
}
