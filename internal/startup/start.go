package startup

import (
	"log"
	"os"

	"github.com/lcaballero/archon2/internal/editor"
	"github.com/lcaballero/archon2/internal/terminal"
)

func Start() {
	logfile, err := os.OpenFile("archon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logfile)
	os.Stdout = logfile
	defer logfile.Close()

	Dump()

	ed, err := editor.NewEditor().Start()
	if err != nil {
		panic(err)
	}
	defer ed.Close()

	ed.AwaitExit()
}

func Dump() {
	log.Println(
		"terminal.ColorDefault",
		terminal.ColorDefault,
		terminal.ColorBlack,
		terminal.ColorRed)
}
