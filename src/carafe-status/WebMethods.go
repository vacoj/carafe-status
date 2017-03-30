package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func webStatus(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&SETTINGS)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(rw, string(blob))
}

func webLanding(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "static/index.html")
}
