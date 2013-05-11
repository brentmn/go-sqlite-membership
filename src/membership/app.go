package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8078", nil)
	if err != nil {
		fmt.Printf("{error: \"%s\"}", err.Error())
	}
}
