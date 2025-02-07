package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/gio"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	for {
		e := <-w.Events()
		fmt.Println(e)
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				fmt.Println("a")
				t := time.Now()
				c := gio.NewContain(gtx, 200.0, 100.0)
				ctx := canvas.NewContext(c)
				if err := canvas.DrawPreview(ctx); err != nil {
					panic(err)
				}
				fmt.Println("b", time.Now().Sub(t))
				return c.Dimensions()
			})
			e.Frame(gtx.Ops)
		}
	}
}
