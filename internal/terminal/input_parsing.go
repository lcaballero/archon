package terminal

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type InputParsing struct{}

func NewInputParsing() *InputParsing {
	return &InputParsing{}
}

func (t *InputParsing) extractEvent(data []byte) (Event, error) {
	event := Event{}

	if len(data) == 0 {
		return event, fmt.Errorf("has 0 length inbuf")
	}

	if data[0] == byte(27) {
		log.Println("escape key pressed")
		event.RuneSize = 0
		event.Key = KeyEsc
		return event, nil
	}

	if data[0] == '\033' {
		n, ok := t.parseEscapeSequence(&event, data)
		log.Printf("n: %d, ok: %t, data:%v", n, ok, data)
		if n != 0 {
			event.RuneSize = n
			return event, nil
		}
	}

	// if we're here, this is not an escape sequence and not an alt sequence
	// so, it's a FUNCTIONAL KEY or a UNICODE character

	// first of all check if it's a functional key
	if Key(data[0]) <= KeySpace || Key(data[0]) == KeyDelete {
		// fill event, pop buffer, return success
		event.Ch = 0
		event.Key = Key(data[0])
		event.RuneSize = 1
		return event, nil
	}

	// the only possible option is utf8 rune
	if r, n := utf8.DecodeRune(data); r != utf8.RuneError {
		event.Ch = r
		event.Key = 0
		event.RuneSize = n
		return event, nil
	}

	return event, nil
}

func (p *InputParsing) parseEscapeSequence(event *Event, buf []byte) (int, bool) {
	bufstr := string(buf)

	if len(bufstr) >= 6 && strings.HasPrefix(bufstr, "\033[M") {
		switch buf[3] & 3 {
		case 0:
			event.Key = MouseLeft
		case 1:
			event.Key = MouseMiddle
		case 2:
			event.Key = MouseRight
		case 3:
			return 6, false
		}
		event.Type = EventMouse // KeyEvent by default

		// wheel up outputs MouseLeft
		if buf[3] == 0x60 || buf[3] == 0x70 {
			event.Key = MouseMiddle
		}

		// the coord is 1,1 for upper left
		event.MouseX = int(buf[4]) - 1 - 32
		event.MouseY = int(buf[5]) - 1 - 32
		return 6, true
	}

	ep := []string{
		"\x1bOP",   // 1
		"\x1bOQ",   // 2
		"\x1bOR",   // 3
		"\x1bOS",   // 4
		"\x1b[15~", // 5
		"\x1b[17~", // 6
		"\x1b[18~", // 7
		"\x1b[19~", // 8
		"\x1b[20~", // 9
		"\x1b[21~", // 10
		"\x1b[23~", // 11
		"\x1b[24~", // 12
		"\x1b[2~",  // 13
		"\x1b[3~",  // 14
		"\x1bOH",   // 15
		"\x1bOF",   // 16
		"\x1b[5~",  // 17
		"\x1b[6~",  // 18
		"\x1bOA",   // 19
		"\x1bOB",   // 20
		"\x1bOD",   // 21
		"\x1bOC",   // 22
	}

	for i, key := range ep {
		if strings.HasPrefix(bufstr, key) {
			event.Ch = 0
			event.Key = Key(0xFFFF - i)
			return len(key), true
		}
	}

	return 0, true
}
