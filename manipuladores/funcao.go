package manipuladores

import (
	"fmt"
	"net/http"
)

//Funcao manipula a funcao
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipular de funcao num pacote")
}
