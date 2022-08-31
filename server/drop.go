package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type drop struct {
	app.Compo
}

func (d *drop) Render() app.UI {
	return app.Div().Body(
		app.Span().Text("Drop File Here Click to upload").Class("drop-zone__prompt"),
		app.Input().Type("file").Name("myFile").Class("drop-zone__input"),
	).Class("drop-zone").
		Style("max-width", "200px").
		Style("height", "200px").
		Style("padding", "25px").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("text-align", "center").
		Style("border", "4px dashed #009578").
		Style("border-radius", "10px").
		Style("font-family", "Consolas").
		Style("font-weight", "500")
}
