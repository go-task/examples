package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed public/*
var publicDir embed.FS

//go:embed templates/*
var templatesDir embed.FS

func main() {
	http.Handle("/public/", withoutCache(http.FileServer(http.FS(publicDir))))
	http.HandleFunc("/", indexPage)

	fmt.Printf("Running web app on :8383")
	if err := http.ListenAndServe(":8383", nil); err != nil {
		panic(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	data, err := templatesDir.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(data); err != nil {
		panic(err)
	}
}

// NOTE: This is fine for development but don't do this on production.
// You'll probably prefer to do cache busting by adding the asset checksum
// as part of the filename or in a query string.
func withoutCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		next.ServeHTTP(w, r)
	})
}
