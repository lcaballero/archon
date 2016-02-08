package grid
import "github.com/lcaballero/archon2/internal/terminal"

type Cell struct {
	Byte byte
	Ch rune
	Width int
	Fg terminal.Attribute
	Bg terminal.Attribute
}

