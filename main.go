package main

import (
	"fmt"
	"net/http"
	"servidor_web/manipuladores"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola mundo!")
	})

	http.HandleFunc("/funcao", manipuladores.Funcao)
	http.HandleFunc("/ola", manipuladores.Ola)
	fmt.Println("Escutando na porta 8181...")
	http.ListenAndServe(":8181", nil)
}
