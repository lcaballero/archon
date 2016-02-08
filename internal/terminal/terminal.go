package terminal

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"syscall"

	"github.com/lcaballero/archon2/internal/term_sys"
)

type Terminal struct {
	termOut          *os.File
	termIn           int // File Descriptor
	originalTermAtts *term_sys.TermIos
	setupTermAtts    *term_sys.TermIos
	width            int
	height           int
	escapes          *EscapeSeqs
	inputParsing     *InputParsing

	onWindowSizeChange chan os.Signal
	onSignalEvent      chan os.Signal

	quitSignals chan *sync.WaitGroup
	quitWriting chan *sync.WaitGroup
}

func NewTerminal() (*Terminal, error) {
	termOut, err := os.OpenFile("/dev/tty", syscall.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}

	termIn, err := syscall.Open("/dev/tty", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	_, err = term_sys.ChangeFileDescriptor(termIn, syscall.F_SETFL, syscall.O_ASYNC|syscall.O_NONBLOCK)
	if err != nil {
		return nil, err
	}

	_, err = term_sys.ChangeFileDescriptor(termIn, syscall.F_SETOWN, syscall.Getpid())
	if runtime.GOOS != "darwin" && err != nil {
		return nil, err
	}

	originalTermAtts, err := term_sys.SetupTerminalAttributes(termOut)
	if err != nil {
		panic(err)
	}

	setupTermAtts, err := term_sys.SetupTerminalAttributes(termOut)
	if err != nil {
		panic(err)
	}

	err = term_sys.SetTerminalAttributes(termOut.Fd(), setupTermAtts)
	if err != nil {
		panic(err)
	}

	w, h, err := term_sys.GetTerminalDims(termOut.Fd())
	if err != nil {
		panic(err)
	}

	t := &Terminal{
		termOut:            termOut,
		termIn:             termIn,
		setupTermAtts:      setupTermAtts,
		originalTermAtts:   originalTermAtts,
		width:              w,
		height:             h,
		escapes:            NewEscapes(),
		onWindowSizeChange: term_sys.NewSignal(syscall.SIGWINCH),
		onSignalEvent:      term_sys.NewSignal(syscall.SIGIO),
		quitSignals:        make(chan *sync.WaitGroup, 0),
		quitWriting:        make(chan *sync.WaitGroup, 0),
		inputParsing:       NewInputParsing(),
	}

	return t, nil
}

//listen starts a providing events over the return InputEvent channel.
func (t *Terminal) forwardEvents() chan Event {
	onInput := make(chan Event, 1000)

	go func() {
		buf := make([]byte, 128)
		for {
			select {
			case <-t.onSignalEvent:
				for {
					n, err := syscall.Read(t.termIn, buf)
					if term_sys.HasInput(err) {
						b := buf[:n]
						event, err := t.inputParsing.extractEvent(b)
						if err != nil {
							log.Println("onSignalEvent", err)
						} else {
							onInput <- event
						}
					}
					if len(t.onSignalEvent) == 0 {
						break
					}
				}

			case <-t.onWindowSizeChange:
				w, h, err := term_sys.GetTerminalDims(t.termOut.Fd())
				if err != nil {
					log.Println(err)
					continue
				}
				ev := Event{Type: EventResize, Width: w, Height: h}
				onInput <- ev

			case wg := <-t.quitSignals:
				log.Println("Quiting Terminal : listen return")
				wg.Done()
				return
			}
		}
	}()

	return onInput
}

// renderWritable
func (t *Terminal) renderWritable() chan io.WriterTo {
	writerTo := make(chan io.WriterTo, 1)

	go func() {
		for {
			select {
			case wg := <-t.quitWriting:
				wg.Done()
				return
			case to := <-writerTo:
				to.WriteTo(t.termOut)
			}
		}
	}()

	return writerTo
}

// Start() starts the terminal sending system notification via the returned channels.
func (t *Terminal) Start() (chan Event, chan io.WriterTo) {
	esc := t.escapes
	t.WriteAll(
		esc.Enter(),
		esc.EnterKeypad(),
		esc.HideCursor(),
		esc.ClearScreen())

	return t.forwardEvents(), t.renderWritable()
}

// write is a convenience function for variable length write commands.
func (t *Terminal) WriteAll(sg ...string) {
	for _, s := range sg {
		t.termOut.WriteString(s)
	}
}

// Close returns the terminal to the original state.
func (t *Terminal) Close() error {
	log.Println("Quiting Terminal")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	t.quitSignals <- wg
	t.quitWriting <- wg
	wg.Wait()

	esc := t.escapes
	t.WriteAll(
		esc.ShowCursor(),
		esc.Sgr0(),
		esc.ClearScreen(),
		esc.Exit(),
		esc.ExitKeypad(),
		esc.MouseExit())

	term_sys.SetTerminalAttributes(t.termOut.Fd(), t.originalTermAtts)
	err := syscall.Close(t.termIn)
	t.termIn = 0 // File descriptor
	if err != nil {
		return fmt.Errorf("Close error %d", err) //TODO: log instead of print
	}
	return nil
}

func (t *Terminal) Dims() (int, int) {
	return t.width, t.height
}
