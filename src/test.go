package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Program crashed")
	}
}
