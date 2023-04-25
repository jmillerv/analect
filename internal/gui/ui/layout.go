package ui

import "fyne.io/fyne/v2"

type FixedSideMenuLayout struct {
	MenuWidth int
}

func (l *FixedSideMenuLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) != 2 {
		return
	}

	menuSize := fyne.NewSize(float32(l.MenuWidth), size.Height)
	contentSize := fyne.NewSize(float32(size.Width)-float32(l.MenuWidth), size.Height)

	objects[0].Resize(menuSize)
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(contentSize)
	objects[1].Move(fyne.NewPos(float32(l.MenuWidth), 0))
}

func (l *FixedSideMenuLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}
