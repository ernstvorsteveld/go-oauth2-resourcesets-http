package main

import (
	"encoding/json"
	"io/ioutil"
)

// FileRoutesLoader loads the routes from file
type FileRoutesLoader struct {
	filename string
}

// LoadRoutes Load the routes
func (f *FileRoutesLoader) LoadRoutes() (*Routes, error) {
	file, re := ioutil.ReadFile(f.filename)
	if re != nil {
		return nil, re
	}
	routes := Routes{}

	ue := json.Unmarshal([]byte(file), &routes)
	if ue != nil {
		return nil, ue
	}
	return &routes, nil
}
