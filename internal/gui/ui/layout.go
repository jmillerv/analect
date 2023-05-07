package ui

import (
	"fyne.io/fyne/v2"
)

type FixedSideMenuLayout struct {
	MenuWidth int
}

func (l *FixedSideMenuLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) != 2 {
		return
	}

	menuSize := fyne.NewSize(200, size.Height)
	contentSize := fyne.NewSize(size.Width-200, size.Height)

	objects[0].Resize(menuSize)
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(contentSize)
	objects[1].Move(fyne.NewPos(200, 0))
}

func (l *FixedSideMenuLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) != 2 {
		return fyne.NewSize(0, 0)
	}

	sideMenuMinSize := objects[0].MinSize()
	contentMinSize := objects[1].MinSize()

	minWidth := float32(l.MenuWidth) + contentMinSize.Width
	minHeight := fyne.Max(sideMenuMinSize.Height, contentMinSize.Height)

	return fyne.NewSize(minWidth, minHeight)
}
