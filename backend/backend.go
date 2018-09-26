package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Hello struct {
	Version string
}

func main() {
	version := os.Getenv("VERSION")
	hello := Hello{version}
	js, _ := json.Marshal(hello)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":3000", nil)
}
