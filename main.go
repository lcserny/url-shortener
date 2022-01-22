package main

import (
	"fmt"
	"net/http"
	"url-shortener/server"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://www.google.com",
		"/yaml-godoc":     "https://www.microsoft.com",
	}
	mapHandler := server.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://www.reddit.com
- path: /urlshort-final
  url: https://www.github.com
`
	yamlHandler, err := server.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on localhost:8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
