package main

import (
	"encoding/csv"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func returnTemplate(w http.ResponseWriter) {
	htmlTemplate, err := os.ReadFile("templates/input.html")
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
	TITLE.Surname = r.FormValue("surname")
	TITLE.Name = r.FormValue("name")
	TITLE.Title = r.FormValue("title")
}

func serveCSV(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/csv")
	// w.Header().Set("Content-Disposition", "attachment; filename=\"title.csv\"")

	writer := csv.NewWriter(w)
	writer.Comma = rune(';')
	defer writer.Flush()

	if err := writer.Write(TITLE.likeSliceOfStrings()); err != nil {
		http.Error(w, "Ошибка при записи CSV", http.StatusInternalServerError)
		return
	}
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}

	var results []string
	for _, title := range titles {
		fullName := title.Surname + " " + title.Name + " " + title.Title
		if strings.Contains(strings.ToLower(fullName), strings.ToLower(query)) {
			results = append(results, fullName)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
