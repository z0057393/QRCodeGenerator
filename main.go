package main

import (
	"QRCode/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/api/generate", controller.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
