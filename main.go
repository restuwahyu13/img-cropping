package main

import (
	"encoding/json"
	_ "image/png"
	"io"
	"log"
	"mime"
	"net/http"
	"slices"
)

type FileUploadRes struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func main() {
	validTipeMime := []string{"image/png", "image/jpeg", "image/jpg"}

	server := http.NewServeMux()

	server.Handle("/", http.FileServer(http.Dir("./")))

	server.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "multipart/form-data")

		if r.Method != http.MethodPost {
			http.Error(w, "Method tidak diperbolehkan", http.StatusMethodNotAllowed)
			return
		}

		file, metadata, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "File input property salah", http.StatusInternalServerError)
			return
		}

		tipeMime, _, err := mime.ParseMediaType(metadata.Header.Get("Content-Type"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if metadata.Size > http.DefaultMaxHeaderBytes {
			http.Error(w, "File terlalu besar", http.StatusBadRequest)
			return
		} else if slices.Index(validTipeMime, tipeMime) == -1 {
			http.Error(w, "Format file tidak valid", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		gambar, err := parsingGambar(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kotak, err := cariKotakHitam(gambar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		potong := potongGambar(gambar, *kotak)
		simpanGambar("output.png", potong)

		w.Header().Add("Content-Type", "application/json")

		res := FileUploadRes{}
		res.Message = "Gambar berhasil di potong"
		res.Url = "http://localhost:8080/output.png"

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Println("HTTP Server berjalan di port: 8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
		return
	}
}
