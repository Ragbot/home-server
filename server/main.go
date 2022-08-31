package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ReadFile(name string) string {
	r_string, _ := os.ReadFile(name)
	return string(r_string)
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static")))

	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.Route("/drop", &drop{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "Hello World example!",
	})
	http.Handle("/drop", &app.Handler{
		Name:        "Drop",
		Description: "File Dropper",
	})

	if err := http.ListenAndServe(":7878", nil); err != nil {
		log.Fatal(err)
	}
}