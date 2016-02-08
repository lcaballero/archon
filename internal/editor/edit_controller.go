package editor

import (
	"io"
	"log"

	"github.com/lcaballero/archon/internal/grid"
	"github.com/lcaballero/archon/internal/sys/terminal"
	"github.com/lcaballero/archon/internal/view"
)


type EditController struct {
	atts       *view.ViewAtts
	onRender   chan io.WriterTo
	exitEditor chan bool
	escapes    *terminal.EscapeSeqs
	fullView   *view.FullView
	grid       *grid.Grid
	cursor     *Cursor
}

func NewEditController(
	atts *view.ViewAtts,
	writer chan io.WriterTo,
	exitEditor chan bool) *EditController {

	cursor := NewCursor(0, 0)
	ec := &EditController{
		atts:       atts,
		onRender:   writer,
		exitEditor: exitEditor,
		escapes:    terminal.NewEscapes(),
		grid:       grid.NewGrid(atts.Dims()),
		fullView:   view.NewFullView("*scratch*", cursor),
		cursor:     cursor,
	}

	ec.fullView.Render(atts, ec.grid)
	writer <- grid.CopyGrid(ec.grid)

	return ec
}

func (d *EditController) Handle(ev terminal.Event) error {
	switch ev.Key {
	case terminal.KeyEnter:
		log.Println("enter pressed")
		d.cursor.Newline()
		d.fullView.Render(d.atts, d.grid)
		d.onRender <- grid.CopyGrid(d.grid)
	case terminal.KeyEsc:
		d.exitEditor <- true
	case terminal.KeyDelete:
		log.Println("terminal.KeyDelete")
		x, y, c := d.handleKey(ev, ev.Key)
		d.grid.Set(x, y, c)
		d.fullView.Render(d.atts, d.grid)
		d.onRender <- grid.CopyGrid(d.grid)
	default:
		x, y, c := d.handleKey(ev, ev.Key)
		d.grid.Set(x, y, c)
		d.fullView.Render(d.atts, d.grid)
		d.onRender <- grid.CopyGrid(d.grid)
	}


	return nil
}

func (d *EditController) handleKey(
	ev terminal.Event, key terminal.Key) (int, int, grid.Cell) {
	b := ev.Byte()
	x, y := d.cursor.Pos()

	// Set Byte
	switch key {
	case terminal.KeySpace, terminal.KeyBackspace:
		b = ' '
		d.cursor.Right()
	case terminal.KeyDelete:
		log.Println("KeyDelete")
		b = ' '
		d.cursor.Left()
		x, y = d.cursor.Pos()   // resets going leftward
	default:
		d.cursor.Right()
	}

	return x, y, grid.Cell{ Byte: b }
}