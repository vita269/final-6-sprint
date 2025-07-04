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

// Корневой хендлер для отображения HTML-формы
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение HTML-шаблона
	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Выполнение шаблона
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}

// Хендлер для обработки загрузки файлов
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Парсинг формы
	err := r.ParseMultipartForm(32 << 20) // 32 MB limit
	if err != nil {
		log.Printf("Ошибка при парсинге формы: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Получение файла из формы
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Ошибка при получении файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение данных файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Конвертация данных
	convertedString, err := service.AutoDetect(string(data))
	if err != nil {
		log.Printf("Ошибка при конвертации: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Генерация имени файла
	timestamp := time.Now().UTC().String()
	ext := filepath.Ext(handler.Filename)
	newFilename := fmt.Sprintf("%s%s", timestamp, ext)

	// Запись результата в локальный файл
	err = os.WriteFile(newFilename, []byte(convertedString), 0644)
	if err != nil {
		log.Printf("Ошибка при записи файла: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	// Возврат результата
	w.Write([]byte(convertedString))
}
