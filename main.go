package main

import (
	"fmt"
	"github.com/h8gi/canvas"
	"math/rand"
)

func main() {
	var x_PosOne float64 = 600
	var x_PosTwo float64 = 300
	var x_PosThree float64 = 800

	var y_PosOne float64 = 800
	var y_PosTwo float64 = 350
	var y_PosThree float64 = 350

	mainCanvas := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     1200,
		Height:    1000,
		FrameRate: 60,
	})

	mainCanvas.Draw(func(ctx *canvas.Context) {
		ctx.DrawPoint(x_PosOne, y_PosOne, 1)
		ctx.DrawPoint(x_PosTwo, y_PosTwo, 1)
		ctx.DrawPoint(x_PosThree, y_PosThree, 1)
		ctx.Fill()

		if ctx.IsMouseDragged {
			var x_PosMain float64 = 370
			var y_PosMain float64 = 450
			for i := 0; i <= 10000; i++ {

				var ball = rand.Intn(3)
				if ball == 0 {
					x_PosMain = (x_PosMain + x_PosOne) / 2
					y_PosMain = (y_PosMain + y_PosOne) / 2
					ctx.DrawPoint(x_PosMain, y_PosMain, 1)
					ctx.Fill()
					fmt.Println(ball)
				} else if ball == 1 {
					x_PosMain = (x_PosMain + x_PosTwo) / 2
					y_PosMain = (y_PosMain + y_PosTwo) / 2
					ctx.DrawPoint(x_PosMain, y_PosMain, 1)
					ctx.Fill()
					fmt.Println(ball)
				} else if ball == 2 {
					x_PosMain = (x_PosMain + x_PosThree) / 2
					y_PosMain = (y_PosMain + y_PosThree) / 2
					ctx.DrawPoint(x_PosMain, y_PosMain, 1)
					ctx.Fill()
					fmt.Println(ball)
				}

			}
		}

	})
}
