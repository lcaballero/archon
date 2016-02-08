package view

import (
	"fmt"

	"github.com/lcaballero/archon/internal/grid"
)

// ClipArea constrains writing bytes to the GridWriter and translated
// to a coordinate object space.
type ClipArea struct {
	grid       grid.GridWriter
	x, y, w, h int
}

// NewClipArea creates a rectangle clipping area that constrains writes
// to a width and height translated to the provided x and y values.
func NewClipArea(grid grid.GridWriter, x, y, w, h int) (*ClipArea, error) {
	if grid == nil {
		return nil, fmt.Errorf("GridWriter provided is nil")
	}
	if x < 0 || y < 0 {
		return nil, fmt.Errorf("Negative ClipArea position x:%d, y:%d", x, y)
	}
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("Negative or zero ClipArea w:%d, h:%d")
	}

	c := &ClipArea{
		grid: grid,
		x:    x,
		y:    y,
		w:    w,
		h:    h,
	}

	return c, nil
}

// Set assigns the byte to the given coordinates relative to
// position and dimensions constrained by the ClipArea.
func (c *ClipArea) Set(x, y int, b grid.Cell) error {
	goodX := 0 <= x && x < c.w
	goodY := 0 <= y && y < c.h
	badPos := !(goodX && goodY)

	if badPos {
		return fmt.Errorf("Bad coordinates for Set() x:%d, y:%d", x, y)
	}

	return c.grid.Set(x+c.x, y+c.y, b)
}
