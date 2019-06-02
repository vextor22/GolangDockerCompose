package restservice

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHelloWorlds(r *mux.Router) {
	r.HandleFunc("/vim", vimgo).Methods("GET")
}
func vimgo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "vim-go")
}
