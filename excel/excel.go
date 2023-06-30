package excel

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func SheetWriteFile(
	credentials []byte,
	token []byte,
	spreadsheetId string,
	spreadsheetIdShow string,
	appCode string,
	newData [][]interface{},
) (err error) {
	ctx := context.Background()
	configGoogleSheet, err := google.ConfigFromJSON(
		credentials,
		"https://www.googleapis.com/auth/spreadsheets",
	)
	if err != nil {
		return
	}
	tok := &oauth2.Token{}
	_ = json.Unmarshal(token, tok)
	client := configGoogleSheet.Client(ctx, tok)

	sheetsService, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return
	}

	readData := appCode
	rangeData := readData + "!A2"
	Spreadsheets := sheetsService.Spreadsheets
	resp, err := Spreadsheets.Values.Get(spreadsheetId, readData).Do()
	if err != nil {
		return
	}

	var newResp [][]interface{}
	newResp = append(newResp, newData...)
	if len(resp.Values) >= 2 {
		newResp = append(newResp, resp.Values[1:]...)
	}
	for i, row := range newResp {
		row[0] = i + 1
		newResp[i] = row
	}

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: newResp,
	})
	_, err = Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Context(ctx).Do()
	_, _ = Spreadsheets.Values.BatchUpdate(spreadsheetIdShow, rb).Context(ctx).Do()
	return err
}
