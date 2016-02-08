// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs syscalls.go
package term_sys

import (
	"os"
	"syscall"
	"unsafe"
)

type TermIos struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed uint32
	Ospeed uint32
}

const (
	syscall_IGNBRK = 0x1
	syscall_BRKINT = 0x2
	syscall_PARMRK = 0x8
	syscall_ISTRIP = 0x20
	syscall_INLCR  = 0x40
	syscall_IGNCR  = 0x80
	syscall_ICRNL  = 0x100
	syscall_IXON   = 0x200
	syscall_OPOST  = 0x1
	syscall_ECHO   = 0x8
	syscall_ECHONL = 0x10
	syscall_ICANON = 0x100
	syscall_ISIG   = 0x80
	syscall_IEXTEN = 0x400
	syscall_CSIZE  = 0x300
	syscall_PARENB = 0x1000
	syscall_CS8    = 0x300
	syscall_VMIN   = 0x10
	syscall_VTIME  = 0x11

	syscall_TCGETS = 0x402c7413
	syscall_TCSETS = 0x802c7414
)

func SetupTerminalAttributes(f *os.File) (*TermIos, error) {
	t, err := GetTerminalAttributes(f.Fd())
	if err != nil {
		return nil, err
	}

	t.Iflag &^= syscall_IGNBRK | syscall_BRKINT | syscall_PARMRK | syscall_ISTRIP | syscall_INLCR | syscall_IGNCR | syscall_ICRNL | syscall_IXON
	t.Oflag &^= syscall_OPOST
	t.Lflag &^= syscall_ECHO | syscall_ECHONL | syscall_ICANON | syscall_ISIG | syscall_IEXTEN
	t.Cflag &^= syscall_CSIZE | syscall_PARENB
	t.Cflag |= syscall_CS8
	t.Cc[syscall_VMIN] = 1
	t.Cc[syscall_VTIME] = 0

	return t, nil
}

// See: http://linux.die.net/man/3/tcgetattr
func GetTerminalAttributes(fd uintptr) (*TermIos, error) {
	termios := &TermIos{}

	r, _, e := syscall.Syscall(
		syscall.SYS_IOCTL, fd,
		uintptr(syscall_TCGETS),
		uintptr(unsafe.Pointer(termios)))

	if r != 0 {
		return nil, os.NewSyscallError("SYS_IOCTL", e)
	}
	return termios, nil
}

func SetTerminalAttributes(fd uintptr, termios *TermIos) error {
	r, _, e := syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(syscall_TCSETS), uintptr(unsafe.Pointer(termios)))
	if r != 0 {
		return os.NewSyscallError("SYS_IOCTL", e)
	}
	return nil
}
