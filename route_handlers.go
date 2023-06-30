package main

import (
	"encoding/json"
	"log"
	"net/http"
	"scopes"
)

// RouteHandlers contains the handlers
type RouteHandlers struct {
	Handlers map[string]*func(w http.ResponseWriter, r *http.Request)
}

// RouteHandler is the set of methods on the handlers available
type RouteHandler interface {
	Add(name string, f *func(w http.ResponseWriter, r *http.Request))
}

// NewRouteHanders creates a map of handlers and names
func NewRouteHanders() RouteHandlers {
	var h = RouteHandlers{
		Handlers: make(map[string]*func(w http.ResponseWriter, r *http.Request)),
	}
	h.Handlers[""] = Add
	return h
}

// Add adds a handler function
func (rh *RouteHandlers) Add(name string, f *func(w http.ResponseWriter, r *http.Request)) {
	rh.Handlers[name] = f
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
