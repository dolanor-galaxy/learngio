package main

import (
	"github.com/gop9/olt/gio/app"
	"github.com/gop9/olt/gio/f32"
	"github.com/gop9/olt/gio/io/pointer"
	"github.com/gop9/olt/gio/io/system"
	"github.com/gop9/olt/gio/layout"
	"github.com/gop9/olt/gio/op"
	"github.com/gop9/olt/gio/op/paint"
	"github.com/gop9/olt/gio/widget/material"
	"image"
	"image/color"
)

var text = "button txt"

// START QUEUE OMIT
func main() {
	go func() {
		w := app.NewWindow()
		button := new(Button)
		gtx := layout.NewContext(w.Queue())
		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				gtx.Reset(e.Config, e.Size)
				button.Layout(gtx)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}

// END QUEUE OMIT

type Button struct {
	pressed bool
}

// START OMIT
func (b *Button) Layout(gtx *layout.Context) {
	for _, e := range gtx.Events(b) { // HLevent
		if e, ok := e.(pointer.Event); ok { // HLevent
			switch e.Type { // HLevent
			case pointer.Press: // HLevent
				b.pressed = true // HLevent
			case pointer.Release: // HLevent
				b.pressed = false // HLevent
			}
		}
	}
	th := material.NewTheme()

	col := color.RGBA{A: 0xff, R: 0xff}
	if b.pressed {
		col = color.RGBA{A: 0xff, G: 0xff}
	}
	pointer.Rect( // HLevent
		image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
	).Add(gtx.Ops) // HLevent
	pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
	drawSquare(gtx.Ops, col)
	th.H6(text).Layout(gtx)
}

// END OMIT

func drawSquare(ops *op.Ops, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{X: 500, Y: 500},
	}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{Rect: square}.Add(ops)
}
