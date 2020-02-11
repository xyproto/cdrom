package cd

import (
	"golang.org/x/sys/unix"
	"syscall"
)

type CD struct {
	fileDescriptor int
}

func New() (*CD, error) {
	fd, err := syscall.Open("/dev/cdrom", syscall.O_RDONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return nil, err
	}
	return &CD{fd}, nil
}

func (c *CD) Eject() {
	unix.IoctlGetInt(c.fileDescriptor, CDROMEJECT)
	unix.IoctlGetInt(c.fileDescriptor, CDROMEJECT_SW)
}

func (c *CD) Close() error {
	return syscall.Close(c.fileDescriptor)
}
