package grid

import (
	"io"
	"github.com/lcaballero/archon/internal/terminal"
	"bytes"
)

type GridWriter interface {
	Set(x, y int, b Cell) error
}

type Grid struct {
	w, h  int
	runes [][]Cell
}

func NewGrid(w, h int) *Grid {
	runes := make([][]Cell, h)
	for i := 0; i < h; i++ {
		runes[i] = make([]Cell, w)
		for j := 0; j < w; j++ {
			runes[i][j] = Cell{Byte:' '}
		}
	}

	g := &Grid{
		w:     w,
		h:     h,
		runes: runes,
	}

	return g
}

//TODO: figure out a way to not need a copy.
func CopyGrid(g *Grid) *Grid {
	rv := &Grid{
		w: g.w,
		h: g.h,
		runes: g.ToWritable(),
	}
	return rv
}

func (g *Grid) Set(x, y int, b Cell) error {
	//TODO: test bounds and return error when out of bounds
	g.runes[y][x] = b
	return nil
}

func (g *Grid) ToWritable() [][]Cell {
	grid := make([][]Cell, g.h)
	for i := 0; i < g.h; i++ {
		curr := g.runes[i]
		cp := make([]Cell, len(curr))
		copy(cp, curr)
		grid[i] = cp
	}
	return grid
}

func (g *Grid) WriteTo(w io.Writer) (int64, error) {
	var total int64 = 0
	fg, bg := terminal.ColorDefault, terminal.ColorDefault

	row := bytes.NewBuffer([]byte{})
	esc := terminal.NewEscapes()

	for i := 0; i < g.h; i++ {
		cellRow := g.runes[i]
		for j := 0; j < g.w; j++ {
			cell := cellRow[j]
			if cell.Fg != fg {
				row.WriteString(esc.Fg(cell.Fg))
				fg = cell.Fg
			}
			if cell.Bg != bg {
				row.WriteString(esc.Bg(cell.Bg))
				bg = cell.Bg
			}
			row.WriteByte(cell.Byte)
		}
		n, _ := w.Write(row.Bytes())
		total += int64(n)
		row.Reset()
	}
	return total, nil
}
