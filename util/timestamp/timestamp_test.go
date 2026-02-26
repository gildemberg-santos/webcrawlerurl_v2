package timestamp_test

import (
	"testing"
	"time"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp_Start(t *testing.T) {
	min := 0.00000001
	max := 0.00001
	timestamp := timestamp.NewTimestamp().Start()
	time.Sleep(1 * time.Microsecond)
	timestamp.End()
	assert.True(t, timestamp.GetTime() >= min)
	assert.True(t, timestamp.GetTime() < max)
}
