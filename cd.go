package cdrom

import (
	"golang.org/x/sys/unix"
	"syscall"
)

type CD struct {
	fileDescriptor int
}

// New opens /dev/cdrom and returns a struct with the file descriptor
func New() (*CD, error) {
	fd, err := syscall.Open("/dev/cdrom", syscall.O_RDONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return nil, err
	}
	return &CD{fileDescriptor: fd}, nil
}

// NewFile opens the given device filename and returns a struct with the file descriptor
func NewFile(deviceFilename string) (*CD, error) {
	fd, err := syscall.Open(deviceFilename, syscall.O_RDONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return nil, err
	}
	return &CD{fileDescriptor: fd}, nil
}

// Ejects uses IOCTL to get CDROMEJECT and CDROMEJECT_SW from the current file descriptor
func (cd *CD) Eject() {
	unix.IoctlGetInt(cd.fileDescriptor, CDROMEJECT)
	unix.IoctlGetInt(cd.fileDescriptor, CDROMEJECT_SW)
}

// Done closes the file descriptor
func (cd *CD) Done() error {
	return syscall.Close(cd.fileDescriptor)
}
