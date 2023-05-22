package main

import (
	"learn-save-file/handler"
	"log"
	"net/http"
)

func main() {
	// Register the file server handler to handle all requests
	http.HandleFunc("/video", handler.VideoHandler)
	http.HandleFunc("/img", handler.ImgHandler)
	http.HandleFunc("/upload", handler.HandleImageUpload)
	log.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)

}
