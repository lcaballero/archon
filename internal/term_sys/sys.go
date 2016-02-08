package term_sys

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func NewSignal(sig syscall.Signal) chan os.Signal {
	c := make(chan os.Signal)
	signal.Notify(c, sig)
	return c
}

func ChangeFileDescriptor(fd int, cmd int, arg int) (val int, err error) {
	r, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	val = int(r)
	if e != 0 {
		err = e
	}
	return val, err
}

func GetTerminalDims(fd uintptr) (int, int, error) {
	var sz struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	winsz := uintptr(syscall.TIOCGWINSZ)
	psz := uintptr(unsafe.Pointer(&sz))
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, winsz, psz)
	if errno != 0 {
		return 0, 0, fmt.Errorf("SYS_IOCTL error %d", errno)
	}
	w, h := int(sz.cols), int(sz.rows)

	return w, h, nil
}

func HasInput(err error) bool {
	hasInput := !(err == syscall.EAGAIN || err == syscall.EWOULDBLOCK)
	return hasInput
}
