package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {

	dataHtml, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "index.html read error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// записываем html данные в тело ответа
	w.Write(dataHtml)
}
func MainHandler(w http.ResponseWriter, r *http.Request) {

	// Парсинг МультиПартФормы
	err := r.ParseMultipartForm(10)
	if err != nil {
		http.Error(w, "ParseMultipartForm error", http.StatusInternalServerError)
		return
	}
	// парсинг файла (name="myFile") из body
	file, fileheader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "FormFile error", http.StatusInternalServerError)
		return
	}
	// отложенное закрытие файла
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "io.ReadAll error", http.StatusInternalServerError)
		return
	}
	translated, err := service.Translate(string(fileBytes))
	if err != nil {
		http.Error(w, "Morse translate error", http.StatusInternalServerError)
		return
	}
	fileName := time.Now().UTC().Format("2025-12-01")
	fileExt := filepath.Ext(fileheader.Filename)
	fileNameRes := fmt.Sprintf("%s%s", fileName, fileExt)
	localFileile, err := os.Create(fileNameRes)
	if err != nil {
		http.Error(w, "local file creat error", http.StatusInternalServerError)
		return
	}
	defer localFileile.Close()
	_, err = localFileile.WriteString(translated)
	if err != nil {
		http.Error(w, "local file write error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(translated))
}
