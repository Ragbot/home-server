package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Gallery struct {
	Id     int      `json:"id"`
	Pages  int      `json:"pages"`
	Images []string `json:"pages_locations"`
	Title  string   `json:"title"`
}

func main() {
	galleries := get_galleries()

	for _, gallery := range galleries {
		fmt.Printf("Id: %d, Pages: %d, ", gallery.Id, gallery.Pages)
		for index, image := range gallery.Images {
			fmt.Printf("Image: %s", image)
			if index != len(gallery.Images) {
				fmt.Print(", ")
			}
		}
	}
}

func get_galleries() []Gallery {
	response, _ := http.Get("http://127.0.0.1:7878/api/galleries/")

	bytes, _ := ioutil.ReadAll(response.Body)
	var galleries []Gallery
	json.Unmarshal(bytes, &galleries)

	return galleries

	// for _, gallery := range galleries {
	// 	fmt.Printf("Id: %d, Pages: %d, ", gallery.Id, gallery.Pages)
	// 	for index, image := range gallery.Images {
	// 		fmt.Printf("Image: %s", image)
	// 		if index != len(gallery.Images) {
	// 			fmt.Print(", ")
	// 		}
	// 	}
	// }
}
