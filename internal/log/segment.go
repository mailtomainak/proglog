package log

type Segment struct {
	store                  *store
	index                  *index
	baseOffset, nextOffset uint64
	config                 Config
}
