package brush

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"zerotomastery.io/pixl/apptype"
)

const (
	Pixel = iota
)

func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, ev *desktop.MouseEvent, x int, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	switch {
	case brush == Pixel:
		pxSize := float32(config.PxSize)
		xOrgin := (float32(x) * pxSize) + config.CanvasOffset.X
		yOrgin := (float32(y) * pxSize) + config.CanvasOffset.Y

		cursorColor := color.NRGBA{80, 80, 80, 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrgin, yOrgin)
		left.Position2 = fyne.NewPos(xOrgin, yOrgin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrgin, yOrgin)
		top.Position2 = fyne.NewPos(xOrgin+pxSize, yOrgin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrgin+pxSize, yOrgin)
		right.Position2 = fyne.NewPos(xOrgin+pxSize, yOrgin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrgin, yOrgin+pxSize)
		right.Position2 = fyne.NewPos(xOrgin+pxSize, yOrgin+pxSize)

		objects = append(objects, left, top, right, bottom)
	}
	return objects
}

func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}
