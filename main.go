package main

import (
	"net/http"

	"github.com/Agatha28/go-crud/controllers/pasiencontroller"
)

func main() {
	http.HandleFunc("/", pasiencontroller.Index)
}
