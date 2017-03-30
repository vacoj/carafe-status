package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func statusWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&SETTINGS)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(rw, string(blob))
}

func webLanding(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(wr, req, "static/index.html")
}
