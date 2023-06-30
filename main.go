package main

import (
	"log"
	"net/http"

	scopes "github.com/ernstvorsteveld/go-oauth2-resourcesets"
)

// HTTPHandler contains the beans
type HTTPHandler struct {
	sdu    scopes.ScopeDescriptionUseCase
	routes *Routes
}

func dispatch() {
	routes, error := configureRoutes()
	if error != nil {
		panic(error)
	}
	var httpHandler = HTTPHandler{
		sdu: scopes.ScopeDescriptionUseCase{
			ScopeDb: scopes.NewInMemoryDB(),
		},
		routes: routes,
	}

	http.HandleFunc("/", httpHandler.IndexHandler)
	http.HandleFunc("/api/", httpHandler.APIHandler)
}

func configureRoutes() (*Routes, error) {
	filename := "./config/routes.json"
	frl := FileRoutesLoader{
		filename: filename,
	}

	routes, error := frl.LoadRoutes()
	if error != nil {
		return nil, error
	}

	for _, element := range routes.Routes {
		http.HandlerFunc(element.Path, getHandler(element.Handler))
	}
	return routes, nil
}

func getHandler(name string) (*func(w http.ResponseWriter, r *http.Request), error) {

}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func main() {
	dispatch()
	log.Fatal(http.ListenAndServe(":9000", nil))
}
