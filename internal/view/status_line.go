package view

import (
	"fmt"
	"strings"

	"github.com/lcaballero/archon/internal/grid"
	"github.com/lcaballero/archon/internal/sys/terminal"
)

type StatusLine struct {
	name   string
	cursor Position
}

func NewStatusLine(name string, c Position) *StatusLine {
	return &StatusLine{
		name:   name,
		cursor: c,
	}
}

func (u *StatusLine) Render(v *ViewAtts, wr grid.GridWriter) error {
	y := v.h - 2

	w, h := u.cursor.Pos()
	pos := fmt.Sprintf("(%d,%d)", w, h)
	help := "[Esc:Quit]"

	textLen := len(u.name) + len(pos) + len(help) + 6
	gap := (v.w - textLen) / 2
	_s := strings.Repeat("-", gap)

	// Name -------- (0,0) ---------- [Esc:Quit]
	status := fmt.Sprintf(" %s %s %s %s %s ", u.name, _s, pos, _s, help)

	str := []byte(status)
	for i, b := range str {
		c := grid.Cell{Byte: b}
		if i < len(str) {
			c.Bg = terminal.ColorWhite
			c.Fg = terminal.ColorBlack
		}
		wr.Set(i, y, c)
	}
	return nil
}
