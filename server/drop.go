package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type drop struct {
	app.Compo
}

func (d *drop) Render() app.UI {
	return app.Div().Body(
		// app.Span().Text("Drop File Here Click to upload").Class("drop-zone__prompt"),
		app.Div().Class("drop-zone__thumb").
			ID("drop-zone__thumb").
			OnClick(d.OnCLick).
			OnDrop(d.OnDrop).
			Style("width", "100%").
			Style("height", "100%").
			Style("border-radius", "10px").
			Style("overflow", "hidden").
			Style("background-color", "#cccccc").
			Style("background-size", "cover").
			Style("position", "relative"),
		app.Input().Type("file").Name("myFile").Class("drop-zone__input").ID("drop_input").
			Style("display", "none"),
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
		Style("font-family", "\"Quicksand\", sans-serif").
		Style("font-weight", "500").
		Style("font-size", "20px").
		Style("color", "#cccccc")
}

func (d *drop) OnCLick(ctx app.Context, e app.Event) {
	e.PreventDefault()

	drop_button := app.Window().GetElementByID("drop_input")
	drop_button.Call("click")
}

func (d *drop) OnDrop(ctx app.Context, e app.Event) {
	e.PreventDefault()

	drop_button := app.Window().GetElementByID("drop_input")

}
