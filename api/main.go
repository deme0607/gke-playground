package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message       string      `json:"message"`
	RequestHeader http.Header `json:"request_header"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myres := Response{Message: "OK", RequestHeader: r.Header}
		jsonStr, _ := json.Marshal(myres)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonStr))
	})

	http.ListenAndServe(":8080", nil)
}
