package day02

import (
	"fmt"
	"strconv"
)

type InvalidID struct {
	seq uint64
	n   int
}

func NewInvalidID(seq uint64, n int) InvalidID {
	return InvalidID{
		seq: seq,
		n:   n,
	}
}

func (id InvalidID) String() string {
	return fmt.Sprintf("%d*%d", id.seq, id.n)
}

func (id InvalidID) Value() (uint64, error) {
	seq := strconv.FormatUint(id.seq, 10)

	value := ""
	for i := 0; i < id.n; i++ {
		value += seq
	}

	return strconv.ParseUint(value, 10, 64)
}
