package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"sheet.report.daily/src/config"
	"sheet.report.daily/src/routers"
)

func main() {
	r := mux.NewRouter()
	routers.Api(r)
	r.HandleFunc("/", Hello)

	http.Handle("/", &config.MuxHandler{ServeMux: r})
	_ = http.ListenAndServe(":8080", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(request.Host))
}
