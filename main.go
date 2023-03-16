package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Server on Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
