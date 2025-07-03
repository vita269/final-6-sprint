package handlers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// корневой сервер для отображения HTML страницы
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// чтение шаблона
	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// выполнение шаблона
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
}

// хендлер для обработки загрузки файлов
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	//парсинг формы
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Ошибка при парсинге формы: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// получение файла
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Ошибка при получении файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// чтение файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// конвертация данных
	convertedString, err := service.AutoDefect(string(data))
	if err != nil {
		log.Printf("Ошибка при конвертации данных: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	//генерация имени файла
	timestamp := time.Now().UTC().String()
	extension := filepath.Ext(handler.Filename)
	newFilename := fmt.Sprintf("%s%s", timestamp, extension)

	// запись в файл
	err = os.WriteFile(newFilename, []byte(convertedString), 0644)
	if err != nil {
		log.Printf("Ошибка при записи файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// возврат результата
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(convertedString))
}
