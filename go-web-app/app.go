package main

import (
	"fmt"
	"net/http"

	"github.com/go-task/examples/go-web-app/boxes/assets"
	"github.com/go-task/examples/go-web-app/boxes/templates"
)

func main() {
	http.Handle("/public/", withoutCache(assets.Handler))
	http.HandleFunc("/", indexPage)

	fmt.Printf("Running web app on :8383")
	if err := http.ListenAndServe(":8383", nil); err != nil {
		panic(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	data, err := templates.ReadFile("templates/index.html")
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
