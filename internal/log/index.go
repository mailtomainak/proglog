package log

import (
	"github.com/tysonmote/gommap"
	"os"
)

const (
	offWidth = 4
	posWidth = 8
	entWidth = offWidth + posWidth
)

type index struct {
	file *os.File
	size uint64
	mmap gommap.MMap
}

func newIndex(f *os.File, c Config) (*index, error) {
	// TODO Avoid compilation error.
	return nil, nil
}

func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	return 0, 0, nil
}
