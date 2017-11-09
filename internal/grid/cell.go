package grid
import "github.com/lcaballero/archon/internal/terminal"

type Cell struct {
	Byte byte
	Ch rune
	Width int
	Fg terminal.Attribute
	Bg terminal.Attribute
}

