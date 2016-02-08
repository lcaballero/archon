package view

import (
	"github.com/lcaballero/archon/internal/grid"
)

type GridWriter interface {
	Set(x, y int, b grid.Cell) error
}

type Renderable interface {
	Render(ViewAtts, GridWriter) error
}

type ViewAtts struct {
	x, y, w, h int
}

func NewViewAtts(x, y, w, h int) *ViewAtts {
	return &ViewAtts{w: w, h: h}
}

func (a *ViewAtts) Dims() (int, int) {
	return a.w, a.h
}
func (a *ViewAtts) Pos() (int, int) {
	return a.x, a.y
}
