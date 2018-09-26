package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Hello struct {
	Version string
}

func main() {
	backend := os.Getenv("BACKEND")

	backendClient := http.Client{
		Timeout: time.Second * 2,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest(http.MethodGet, backend, nil)
		req.Header.Set("Content", "Application/json")
		resp, err := backendClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		hello := Hello{}
		jsonErr := json.Unmarshal(body, &hello)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.Write([]byte(hello.Version))

	})

	http.ListenAndServe(":3001", nil)
}
