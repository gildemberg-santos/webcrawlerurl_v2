package pkg

import (
	"time"
)

type Timestamp struct {
	TimeStart int64
	TimeEnd   int64
}

func (t *Timestamp) Start() {
	t.TimeStart = time.Now().UnixNano()
}

func (t *Timestamp) End() {
	t.TimeEnd = time.Now().UnixNano()
}

func (t *Timestamp) GetTime() float64 {
	return (float64(t.TimeEnd) - float64(t.TimeStart)) / 1000000000.0
}
