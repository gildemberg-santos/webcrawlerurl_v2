package timestamp

import (
	"time"
)

type Timestamp struct {
	TimeStart int64
	TimeEnd   int64
}

func NewTimestamp() *Timestamp {
	return &Timestamp{}
}

func (t *Timestamp) Start() *Timestamp {
	t.TimeStart = time.Now().UnixNano()
	return t
}

func (t *Timestamp) End() *Timestamp {
	t.TimeEnd = time.Now().UnixNano()
	return t
}

func (t *Timestamp) GetTime() float64 {
	return (float64(t.TimeEnd) - float64(t.TimeStart)) / 1000000000.0
}
