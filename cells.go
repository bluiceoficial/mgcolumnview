// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package mgcolumnview

import "fyne.io/fyne/v2"

// layout personalizado com colunas fixas
type fixedColumnsLayout struct {
	colWidths []float32
}

func (l *fixedColumnsLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	x := float32(0)
	for i, o := range objects {
		if i >= len(l.colWidths) {
			break
		}
		w := l.colWidths[i]
		o.Resize(fyne.NewSize(w, size.Height))
		o.Move(fyne.NewPos(x, 0))
		x += w
	}
}

func (l *fixedColumnsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var totalWidth float32
	var maxHeight float32
	for i, o := range objects {
		if i >= len(l.colWidths) {
			break
		}
		totalWidth += l.colWidths[i]
		h := o.MinSize().Height
		if h > maxHeight {
			maxHeight = h
		}
	}
	return fyne.NewSize(totalWidth, maxHeight)
}
