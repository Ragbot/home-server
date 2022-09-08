package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type drop struct {
	app.Compo
}

func (d *drop) Render() app.UI {
	return app.Div().Body(
		NavBar(),
		TopBar(),
		FileList(),
	)
}

func NavBar() app.HTMLDiv {
	// Nav-Side Bar
	return app.Div().ID("nav-bar").Body(
		// Button List
		app.Ul().Body(
			app.P().Text("Close").
				OnClick(func(ctx app.Context, e app.Event) {
					nav_bar := app.Window().GetElementByID("nav-bar")
					nav_bar.Get("style").Set("display", "none")
				}).
				Style("font-weight", "bold").
				Style("user-select", "none"),
			app.P().Text("Home").OnClick(func(ctx app.Context, e app.Event) {
				app.Window().Get("location").Call("assign", "/")
			}), // Home Page Button
			app.P().Text("Drop"), // Drop Page Button
		).
			// Button List Styles
			Styles(map[string]string{
				"list-style-type": "none",
				"margin":          "0",
				"padding":         "0",
				"z-index":         "1",
				"font-family":     "Consolas",
				"user-select":     "none",
				"weight":          "500",
			}),
	).
		// Nav Bar Styles
		Styles(map[string]string{
			"height":     "100%",
			"width":      "200px",
			"position":   "fixed",
			"z-index":    "1",
			"overflow":   "auto",
			"background": "white",
		})
}

func TopBar() app.HTMLDiv {
	return app.Div().Class("tob-bar").Body(
		app.Img().Src("web/neon.png").
			Styles(map[string]string{
				"height": "20%",
				"width":  "20%",
			}),
	). // Top Bar Styles
		Styles(map[string]string{
			"float":  "right",
			"height": "25%",
		}).
		OnClick(func(ctx app.Context, e app.Event) {
			nav_bar := app.Window().GetElementByID("nav-bar")
			nav_bar.Get("style").Set("display", "flex")
		})
}

func FileList() app.HTMLDiv {
	return app.Div().Class("file-list").Body(
		app.P().Text("FILE"),
	).Styles(map[string]string{
		"font-weight": "bold",
		"font-family": "consolas",
	})
}
