package main

import (
	"net/http"
)

func startWebApp() {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		// landing page
		"/":       webLanding,
		"/status": webStatus,
	}

	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// static content does not require authentication
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Host server
	panic(http.ListenAndServe(":"+SETTINGS.WebPort, nil))
}
