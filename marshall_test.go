package main

import (
	"net/url"
	"testing"
)

func Test_should_marshall_and_unmarshall(t *testing.T) {
	// var s = scopes.NewScope("View", GetURL("http://www.example.com/icons/view.png"))
	// var sb, e = json.Marshal(s)
	// if e != nil {
	// 	t.Errorf("We have a marshalling issue")
	// }

	// fmt.Printf("We have bytes: %s", string(sb))
}

func GetURL(s string) url.URL {
	url, _ := url.Parse(s)
	return *url
}
