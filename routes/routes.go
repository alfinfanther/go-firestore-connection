package routes

import (
	"gofirestore/products"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getIndex)
	router.Handle("/products", http.HandlerFunc(products.ProductAll)).Methods("GET")

	return router
}
func getIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		log.Fatal("Error template rendering", err)
	}
	t.Execute(w, nil)
}
