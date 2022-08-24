package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xhfmvls/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}