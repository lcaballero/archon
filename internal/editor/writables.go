package editor

import "io"

type WritableString string

func (w WritableString) WriteTo(out io.Writer) (int64, error) {
	n, err := out.Write([]byte(w))
	return int64(n), err
}

type GridWritable [][]byte

func (g GridWritable) WriteTo(w io.Writer) (int64, error) {
	var total int64 = 0
	h := len(g)
	for i := 0; i < h; i++ {
		n, err := w.Write(g[i])
		if err != nil {
			return total, err
		}
		total += int64(n)
	}
	return total, nil
}
