package handler

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

func HandleImageUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve the uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil || format != "jpeg" {
		http.Error(w, "Invalid or unsupported image format", http.StatusBadRequest)
		return
	}
	uuid := uuid.New()
	fileLocation := "img/" + uuid.String() + ".jpg"
	outputFile, err := os.Create(fileLocation)
	if err != nil {
		http.Error(w, "Failed to create output file"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Image saved as", outputFile.Name())
}

// Create a file handler function
func VideoHandler(w http.ResponseWriter, r *http.Request) {
	// Open the file

	filePath := "video/rickroll.mp4"
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Set the Content-Type header based on the file's MIME type
	contentType, _ := mimetype.DetectFile(filePath) // Default MIME type
	// You can use a package like "github.com/gabriel-vasile/mimetype" to detect the MIME type based on file extension or content.
	// For example: contentType, _ = mimetype.DetectFile(filePath)

	w.Header().Set("Content-Type", contentType.String())
	filename := filepath.Base(filePath)
	w.Header().Set("Content-Disposition", "inline; filename="+filename)

	// Serve the file content
	http.ServeContent(w, r, "", fileInfo.ModTime(), file)
}
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	// Open the file

	filePath := "img/fivenightsjumpscare.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Set the Content-Type header based on the file's MIME type
	contentType, _ := mimetype.DetectFile(filePath) // Default MIME type
	// You can use a package like "github.com/gabriel-vasile/mimetype" to detect the MIME type based on file extension or content.
	// For example: contentType, _ = mimetype.DetectFile(filePath)

	w.Header().Set("Content-Type", contentType.String())
	filename := filepath.Base(filePath)
	w.Header().Set("Content-Disposition", "inline; filename="+filename)

	// Serve the file content
	http.ServeContent(w, r, "", fileInfo.ModTime(), file)
}
