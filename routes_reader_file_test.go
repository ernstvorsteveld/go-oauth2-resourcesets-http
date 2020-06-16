package main

import (
	"fmt"
	"testing"
)

func Test_should_read_routes_from_file(t *testing.T) {
	filename := "./config/routes.json"
	frl := FileRoutesLoader{
		filename: filename,
	}

	routes, error := frl.LoadRoutes()
	if error != nil {
		t.Errorf("Should be able to read file %s\n", filename)
	}
	if len(routes.Routes) != 2 {
		t.Errorf("The number of routes found is not equals to 2\n")
	}

	for i := 0; i < len(routes.Routes); i++ {
		fmt.Printf("Route %s for methods %v and handler %s\n", routes.Routes[i].Path, routes.Routes[i].Methods, routes.Routes[i].Handler)
	}

}
