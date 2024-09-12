package main

import (
	"fmt"
	"net/http"
	"os"
)

var imageTag string

func main() {
	imageTag = os.Getenv("IMAGE_TAG")
	if imageTag == "" {
		imageTag = "unknown"
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/tag", handleTag)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! Visit /tag to see the image tag.")
}

func handleTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Image tag: %s", imageTag)
}