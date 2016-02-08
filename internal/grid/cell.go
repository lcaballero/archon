package grid
import "github.com/lcaballero/archon/internal/sys/terminal"

type Cell struct {
	Byte byte
	Ch rune
	Width int
	Fg terminal.Attribute
	Bg terminal.Attribute
}

