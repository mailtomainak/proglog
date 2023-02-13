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
