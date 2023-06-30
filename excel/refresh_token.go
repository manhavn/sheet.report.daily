package excel

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type GoogleVerificationsToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

type GoogleVerificationsRefreshToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type ClientSecretGoogleVerifications struct {
	Web struct {
		ClientId                string   `json:"client_id"`
		ProjectId               string   `json:"project_id"`
		AuthUri                 string   `json:"auth_uri"`
		TokenUri                string   `json:"token_uri"`
		AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
	} `json:"web"`
}

func RefreshTokenSheet(
	gvt *GoogleVerificationsToken,
	credentials []byte,
	fileName string,
) (newDataGvt []byte) {
	var configJson ClientSecretGoogleVerifications
	_ = json.Unmarshal(credentials, &configJson)
	checkTokenTimeout := true
	if gvt.Expiry != "" {
		checkTime, err := time.Parse(time.RFC3339, gvt.Expiry)
		if err == nil {
			checkTokenTimeout = time.Until(checkTime.Add(time.Minute*-15)) < 0
		}
	}
	var checkRefreshToken bool
	if configJson.Web.ClientId != "" && checkTokenTimeout {
		req, err := http.NewRequest(
			"POST",
			"https://www.googleapis.com/oauth2/v4/token",
			strings.NewReader(url.Values{
				"client_id":     {configJson.Web.ClientId},
				"client_secret": {configJson.Web.ClientSecret},
				"refresh_token": {gvt.RefreshToken},
				"grant_type":    {"refresh_token"},
			}.Encode()),
		)
		if err != nil {
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		var resp *http.Response
		client := http.DefaultClient
		resp, err = client.Do(req)

		if err != nil {
			return
		}

		defer resp.Body.Close()
		status := resp.StatusCode
		responseBody, _ := io.ReadAll(resp.Body)

		if status == 200 {
			var rft GoogleVerificationsRefreshToken
			_ = json.Unmarshal(responseBody, &rft)
			if rft.ExpiresIn > 0 {
				newExpiry := time.Now().
					Add(time.Second * time.Duration(rft.ExpiresIn)).
					UTC().
					Format(time.RFC3339)
				(*gvt).AccessToken = rft.AccessToken
				(*gvt).TokenType = rft.TokenType
				(*gvt).Expiry = newExpiry
				checkRefreshToken = true
			}
		}
	}
	newDataGvt, _ = json.Marshal(gvt)
	if checkRefreshToken {
		_ = os.WriteFile(fileName, newDataGvt, 0o755)
	}
	return
}
