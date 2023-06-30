# sheet.report.daily

config `sheet_connect/credentials.json`

```json
{
  "web": {
    "client_id": "____.apps.googleusercontent.com",
    "project_id": "_______",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_secret": "_______",
    "redirect_uris": ["http://localhost:8080"]
  }
}
```

connect sheet: auto generator `sheet_connect/token.json`

```shell
 cd sheet_connect
 go run quickstart.go

 # http://localhost:8080/?state=state-token&code=4%2F0AZEOvhUa5tz7v8oAryIpPyXfZGwqesyfVzfzWmCd-AjFEPkVftQ7AImMyU4-6PJgE4zLzw&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fspreadsheets
 # COPY 4%2F0AZEOvhUa5tz7v8oAryIpPyXfZGwqesyfVzfzWmCd-AjFEPkVftQ7AImMyU4-6PJgE4zLzw
 # Replace %2F = /
 # Paste to
 # Go to the following link in your browser then type the authorization code:
 # https://....................
 # 4/0AZEOvhUa5tz7v8oAryIpPyXfZGwqesyfVzfzWmCd-AjFEPkVftQ7AImMyU4-6PJgE4zLzw
```

api-test.http

```http request
POST http://localhost:8080/rewrite-excel-write-sheet
Content-Type: application/x-www-form-urlencoded

custom_month=2023-06

###
GET http://localhost:8080/cronjob-excel-write-sheet
Accept: application/json

###
POST http://localhost:8080/submit-form-data
Content-Type: application/x-www-form-urlencoded

full_name=Test&email=test@gmail.com&phone=0987654321

###
GET http://localhost:8080/migration-create-table
Accept: application/json

###
POST http://localhost:8080/cronjob/set-api-start
Content-Type: application/x-www-form-urlencoded

api=aHR0cDovL2xvY2FsaG9zdDo4MDgwL2Nyb25qb2ItZXhjZWwtd3JpdGUtc2hlZXQ=&start=true&spec=*_*_*_*_*

###
```
