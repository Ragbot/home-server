package main

type Gallery struct {
	Id     int      `json:"id"`
	Pages  int      `json:"pages"`
	Images []string `json:"pages_locations"`
	Title  string   `json:"title"`
}
