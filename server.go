package main

import (
	"html/template"
	"net/http"
	"os"
)

func returnTemplate(w http.ResponseWriter) {
	htmlTemplate, err := os.ReadFile("template/input.html")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("input").Parse(string(htmlTemplate))
	if err != nil {
		http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при рендеринге шаблона", http.StatusInternalServerError)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	// Убедимся, что метод запроса - POST
	if r.Method != http.MethodPost {
		returnTemplate(w)
	}
	r.ParseForm()
	title := r.FormValue("title")
	file, err := os.OpenFile("csv/title.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Ошибка при открытии файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err := os.WriteFile("csv/title.csv", []byte(title+"\n"), 0644); err != nil {
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}
}

func sendCommand(w http.ResponseWriter, r *http.Request) {
	goTitle()
}
