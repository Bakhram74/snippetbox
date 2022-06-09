package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("address", ":4000", "Сетевой адрес HTTP")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Запуск сервера на %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
