package main

import (
	"fmt" // Standard
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func ReadFile(name string) string {
	r_string, _ := os.ReadFile(name)
	return string(r_string)
}

func WriteFile(file string, data string) {
	os.WriteFile(file, []byte(data), os.ModeIrregular)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

var ver string

func main() {
	ver = ReadFile("version")
	version, _ := strconv.ParseFloat(ver, 32)
	version = roundFloat(version, 2)
	fmt.Println(version)
	WriteFile("version", fmt.Sprint(roundFloat((version+0.01), 2)))
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
		Title:       "Drop",
	})

	http.ListenAndServe(":8000", nil)
}
