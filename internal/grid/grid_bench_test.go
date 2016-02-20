package grid


import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)


func TestGrid(t *testing.T) {
	Convey("Just a blank test", t, func() {
		So(false, ShouldBeFalse)
	})
}

var grid_value *Grid = nil

func Benchmark_CopyGrid(b *testing.B) {
	base_grid := NewGrid(100, 100)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grid_value = CopyGrid(base_grid)
	}
}
