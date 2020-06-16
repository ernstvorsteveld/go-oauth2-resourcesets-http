package main

// Routes the list of routes
type Routes struct {
	Routes []Route `json:"routes"`
}

// Route is the path, http method and handler combination
type Route struct {
	Path    string   `json: "path"`
	Methods []string `json: "methods"`
	Handler string   `json: "handler"`
}

// RoutesLoader loads the routes
type RoutesLoader interface {
	LoadRoutes() (*Routes, error)
}
