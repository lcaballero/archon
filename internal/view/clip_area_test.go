package view

import (
	"testing"

	"github.com/lcaballero/archon/internal/grid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClipArea(t *testing.T) {

	Convey("Set byte to negative y", t, func() {
		g := &grid.Grid{}
		c, _ := NewClipArea(g, 0, 0, 10, 10)
		err := c.Set(0, -1, 'b')
		So(err, ShouldNotBeNil)
	})

	Convey("Set byte to negative x", t, func() {
		g := &grid.Grid{}
		c, _ := NewClipArea(g, 0, 0, 10, 10)
		err := c.Set(-1, 0, 'b')
		So(err, ShouldNotBeNil)
	})

	Convey("New ClipArea should not allow negative 0 width", t, func() {
		g := &grid.Grid{}
		c, err := NewClipArea(g, 0, 0, 0, 10)
		So(c, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("New ClipArea should not allow negative Y position", t, func() {
		g := &grid.Grid{}
		c, err := NewClipArea(g, 0, -1, 10, 10)
		So(c, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("New ClipArea should not allow negative X position", t, func() {
		g := &grid.Grid{}
		c, err := NewClipArea(g, -1, 0, 10, 10)
		So(c, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("New ClipArea should not allow nil GridWriter", t, func() {
		g, err := NewClipArea(nil, 0, 0, 10, 10)
		So(g, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}
