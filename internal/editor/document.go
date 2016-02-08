package editor

import (
	"github.com/lcaballero/archon2/internal/grid"
	"github.com/lcaballero/archon2/internal/view"
)

type Document struct {
	name      string // ex: document.go
	extension string // ex: .go

	cursor   interface{} // position of cursor in doc
	theme    interface{} // mapping of type to colors
	viewAtts interface{} // view x,y,w,h

	// Text data as byte written for writing to a GridWriter.
	buffer interface{}
}

func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Render(v *view.ViewAtts, wr grid.GridWriter) error {
	return nil
}
