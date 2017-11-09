package editor

import (
	"io"
	"log"
	"sync"

	"github.com/lcaballero/archon/internal/terminal"
	"github.com/lcaballero/archon/internal/view"
)

type Editor struct {
	term           *terminal.Terminal
	onInput        chan terminal.Event
	onRender       chan io.WriterTo
	quitInput      chan *sync.WaitGroup
	exitEditor     chan bool
	editController *EditController
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) listen() {
	go func() {
		for {
			select {
			case wg := <-e.quitInput:
				log.Println("Quitting Editor : listen return")
				wg.Done()
				return

			case ev := <-e.onInput:
				e.editController.Handle(ev)
			}
		}
	}()
}

func (e *Editor) AwaitExit() {
	for {
		select {
		case <-e.exitEditor:
			log.Println("Awaiting Close : listen return")
			return
		}
	}
}

func (e *Editor) Start() (*Editor, error) {
	term, err := terminal.NewTerminal()
	if err != nil {
		return nil, err
	}

	w, h := term.Dims()

	e.term = term
	e.quitInput = make(chan *sync.WaitGroup, 0)
	e.exitEditor = make(chan bool, 0)
	e.onInput, e.onRender = e.term.Start()

	e.editController = NewEditController(
		view.NewViewAtts(0, 0, w, h),
		e.onRender,
		e.exitEditor,
	)
	e.listen()

	return e, nil
}

func (e *Editor) Close() error {
	log.Println("Quitting Editor")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	e.quitInput <- wg
	wg.Wait()
	return e.term.Close()
}
