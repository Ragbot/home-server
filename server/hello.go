package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo

	text string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Img().Src("/web/neon.png").Style("width", "100%"),
		app.H1().Body(
			app.If(h.text != "",
				app.Text(h.text),
			).Else(
				app.Text("Hello World!"),
			),
		).OnClick(h.OnClick),
		h.Book(Gallery{
			Id:     0,
			Pages:  0,
			Images: []string{"https://github.com/maxence-charriere/go-app/raw/master/docs/web/images/go-app.png"},
		}),
	).
		Style("margin", "auto").
		Style("width", "50%").
		Style("padding", "10px").
		Style("user-select", "none")
}

func (h *hello) Book(gallery Gallery) app.HTMLDiv {
	return app.Div().Body(
		app.Img().Src(gallery.Images[0]).
			Style("width", "25%"),
		app.P().Text(fmt.Sprint(gallery.Id)),
	).Class("gallery-container")
}

func (h *hello) OnClick(ctx app.Context, e app.Event) {
	if h.text == "Clicked" {
		h.text = ""
	} else {
		h.text = "Clicked"
	}
}
