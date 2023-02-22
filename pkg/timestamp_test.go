package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp_Start(t *testing.T) {
	timestamp := Timestamp{}
	timestamp.Start()
	time.Sleep(1 * time.Second)
	timestamp.End()
	assert.True(t, timestamp.GetTime() >= 1.0)
	assert.True(t, timestamp.GetTime() < 2.0)
}
