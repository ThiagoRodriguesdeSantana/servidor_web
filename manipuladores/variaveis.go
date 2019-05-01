package manipuladores

import "html/template"

//Modelos armazenao modelos de pagina
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
