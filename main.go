package main

import (
	"fmt"
	"log"
	"net/http"
	"searcher/pkg/searcher"
)

func init() {
	err := searcher.ProcessFile() //обработка файлов и заполнение мапы
	if err != nil {
		log.Fatal("Failed initiation", err)
	}
}

func main() {
	http.HandleFunc("/files/search/", searcher.Search)

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
