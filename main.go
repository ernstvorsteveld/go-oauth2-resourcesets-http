package main

import (
	"encoding/json"
	"log"
	"net/http"

	scopes "github.com/ernstvorsteveld/go-oauth2-resourcesets"
)

// HTTPHandler contains the beans
type HTTPHandler struct {
	sdu scopes.ScopeDescriptionUseCase
}

// IndexHandler handles index.html
func (h *HTTPHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index handler")
	w.Write([]byte("index handler\n"))
}

// APIHandler handles all REST Api calls
func (h *HTTPHandler) APIHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API handler")
	if r.Method == "GET" {
		// scope, error := h.sdu.Get(*r.URL)
		// if error != nil {
		// 	http.Error(w, "Scope does not exist", http.StatusNotFound)
		// }
		// js, _ := json.Marshal(scope)
		// w.Write(js)
	} else if r.Method == "POST" {
		s := scopes.Scope{}
		json.NewDecoder(r.Body).Decode(&s)
		w.Write([]byte("POST on /api\n"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func dispatch() {
	var httpHandler = HTTPHandler{
		sdu: scopes.ScopeDescriptionUseCase{
			ScopeDb: scopes.NewInMemoryDB(),
		},
	}
	http.HandleFunc("/", httpHandler.IndexHandler)
	http.HandleFunc("/api/", httpHandler.APIHandler)
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func main() {
	dispatch()
	log.Fatal(http.ListenAndServe(":9000", nil))
}
