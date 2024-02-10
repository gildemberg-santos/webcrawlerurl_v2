package timestamp_test

import (
	"testing"
	"time"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp_Start(t *testing.T) {
	timestamp := timestamp.NewTimestamp().Start()
	time.Sleep(1 * time.Second)
	timestamp.End()
	assert.True(t, timestamp.GetTime() >= 1.0)
	assert.True(t, timestamp.GetTime() < 2.0)
}
