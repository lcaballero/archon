package editor

type Cursor struct {
	x, y int
}

func NewCursor(x, y int) *Cursor {
	return &Cursor{x: x, y: y}
}

func (c *Cursor) Pos() (int, int) {
	return c.x, c.y
}

func (c *Cursor) Right() {
	c.x++
}
func (c *Cursor) Left() {
	c.x--
}
func (c *Cursor) Down() {
	c.y++
}
func (c *Cursor) Up() {
	c.y--
}
func (c *Cursor) Newline() {
	c.Down()
	c.x = 0
}
