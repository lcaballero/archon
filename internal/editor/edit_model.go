package editor

import "github.com/lcaballero/archon2/internal/view"

type EditModel struct {
	documents []interface{}  // all open buffers
	viewAtts  *view.ViewAtts // dimensions of the terminal window
}

func NewEditModel() *EditModel {
	return &EditModel{}
}
