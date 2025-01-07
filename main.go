package main

import (
	"fmt"
	"net/http"
)

var TITLE = LowerThird{
	Surname: "Лавров",
	Name:    "Сергей Викторович",
	Title:   "Министр иностранных дел Российской федерации",
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/title.csv", serveCSV)

	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
