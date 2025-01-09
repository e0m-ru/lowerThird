package main

import (
	"fmt"
	"net/http"
)

// main initializes the HTTP server, sets up the route handlers, and starts
// listening on port 8080. It handles two routes: the root route ("/") which
// serves HTML content, and the "/title.csv" route which serves CSV content.
// If the server fails to start, it logs the error to the console.
func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/titles.csv", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/titles.csv")
	})
	http.HandleFunc("/title.csv", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/title.csv")
	})

	fmt.Println("Сервер запущен на :3456")
	if err := http.ListenAndServe(":3456", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
