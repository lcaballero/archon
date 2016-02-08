package terminal

import (
	"fmt"
)

type (
	OutputMode int
	Attribute  uint16
)

const (
	seq_enter_ca     = iota // 1
	seq_exit_ca             // 2
	seq_show_cursor         // 3
	seq_hide_cursor         // 4
	seq_clear_screen        // 5
	seq_sgr0                // 6
	seq_underline           // 7
	seq_bold                // 8
	seq_blink               // 9
	seq_reverse             // 10
	seq_enter_keypad        // 11
	seq_exit_keypad         // 12
	seq_enter_mouse         // 13
	seq_exit_mouse          // 14

	// Don't have escape sequences for these
	seq_max_funcs // 15
)

// Output mode. See SetOutputMode function.
//go:generate stringer -type=OutputMode
const (
	OutputCurrent OutputMode = iota
	OutputNormal
	Output256
	Output216
	OutputGrayscale
)

const (
	ColorDefault Attribute = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	AttrBold
	AttrUnderline
	AttrReverse
)

type EscapeSeqs struct {
	Escapes []string
}

func NewEscapes() *EscapeSeqs {
	esc := &EscapeSeqs{
		Escapes: []string{
			"\x1b[?1049h",        // 1      Enter
			"\x1b[?1049l",        // 2      Exit
			"\x1b[?12l\x1b[?25h", // 3      ShowCursor
			"\x1b[?25l",          // 4      HideCursor
			"\x1b[H\x1b[2J",      // 5      ClearScreen
			"\x1b(B\x1b[m",       // 6      Sgr0 (Select Graphic Rendition)
			"\x1b[4m",            // 7      Underline
			"\x1b[1m",            // 8      Bold
			"\x1b[5m",            // 9      Blink
			"\x1b[7m",            // 10     Reverse
			"\x1b[?1h\x1b=",      // 11     EnterKeyPad
			"\x1b[?1l\x1b>",      // 12     ExitKeyPad
			"\x1b[?1000h",        // 13     MouseEnter
			"\x1b[?1000l",        // 14     MouseExit
		},
	}

	return esc
}

func (esc *EscapeSeqs) Enter() string {
	return esc.Escapes[seq_enter_ca]
}
func (esc *EscapeSeqs) Exit() string {
	return esc.Escapes[seq_exit_ca]
}
func (esc *EscapeSeqs) ShowCursor() string {
	return esc.Escapes[seq_show_cursor]
}
func (esc *EscapeSeqs) HideCursor() string {
	return esc.Escapes[seq_hide_cursor]
}
func (esc *EscapeSeqs) ClearScreen() string {
	return esc.Escapes[seq_clear_screen]
}
func (esc *EscapeSeqs) Sgr0() string {
	return esc.Escapes[seq_sgr0]
}
func (esc *EscapeSeqs) Underline() string {
	return esc.Escapes[seq_underline]
}
func (esc *EscapeSeqs) Bold() string {
	return esc.Escapes[seq_bold]
}
func (esc *EscapeSeqs) Blink() string {
	return esc.Escapes[seq_blink]
}
func (esc *EscapeSeqs) Reverse() string {
	return esc.Escapes[seq_reverse]
}
func (esc *EscapeSeqs) EnterKeypad() string {
	return esc.Escapes[seq_enter_keypad]
}
func (esc *EscapeSeqs) ExitKeypad() string {
	return esc.Escapes[seq_exit_keypad]
}
func (esc *EscapeSeqs) MouseEnter() string {
	return esc.Escapes[seq_enter_mouse]
}
func (esc *EscapeSeqs) MouseExit() string {
	return esc.Escapes[seq_exit_mouse]
}
func (esc *EscapeSeqs) Fg(a Attribute) string {
	//	return fmt.Sprintf("%s%d%s", "\033[38;5;", uint64(a-1), "m")
	return fmt.Sprintf("%s%d%s", "\x1b[;", uint64(a-1), "m")
}
func (esc *EscapeSeqs) Bg(a Attribute) string {
	//	return fmt.Sprintf("%s%d%s", "\033[48;5;", uint64(a-1), "m")
	return fmt.Sprintf("%s%d%s", "\x1b[;", uint64(a-1), "m")
}
