package main

import (
	"flag"
	"fmt"
	"net/http"
)

func init() {
	flag.Parse()
	fmt.Printf("A: %v\n", *address)
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/send-command", sendCommand)
	http.HandleFunc(
		"/titles.csv",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "csv/titles.csv")
		})
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.HandleFunc("/title.csv", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "csv/title.csv")
	})
	fmt.Println("Сервер запущен на :3456")
	if err := http.ListenAndServe(":3456", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
