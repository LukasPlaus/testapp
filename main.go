package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Print("Starting webappmain...")

	appVersion := getVersion()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! Application version: %s\n\n", appVersion)
		fmt.Fprintf(w, "Client ip:port - %s\n", r.RemoteAddr)

		fmt.Printf("Someone reached us from IP:%s\n ", r.RemoteAddr)
	})

	fmt.Print(" started.")
	http.ListenAndServe(":38181", nil)
}

func getVersion() string {
	releaseTagFile := "release_version.txt"

	dat, err := ioutil.ReadFile(releaseTagFile)
	if err != nil {
		fmt.Printf("Filed reading  %s.", releaseTagFile)
		return "unknown"
	}

	return string(dat)
}
