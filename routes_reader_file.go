package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// FileRoutesLoader loads the routes from file
type FileRoutesLoader struct {
	filename string
}

// LoadRoutes Load the routes
func (f *FileRoutesLoader) LoadRoutes() (*Routes, error) {
	log.Printf("Loading routes configuration from file %s\n", f.filename)
	file, re := ioutil.ReadFile(f.filename)
	if re != nil {
		return nil, re
	}
	routes := Routes{}

	ue := json.Unmarshal([]byte(file), &routes)
	if ue != nil {
		return nil, ue
	}
	logRoutesConfiguration(&routes)
	return &routes, nil
}

func logRoutesConfiguration(r *Routes) {
	for _, route := range r.Routes {
		log.Printf("Route, path %s, methods %v, handler %s.\n", route.Path, route.Methods, route.Handler)
	}
}
