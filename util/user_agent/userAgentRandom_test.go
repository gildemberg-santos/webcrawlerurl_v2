package useragent_test

import (
	"testing"

	useragent "github.com/gildemberg-santos/webcrawlerurl_v2/util/user_agent"
	"github.com/stretchr/testify/assert"
)

func TestUserAgentRandom_Call(t *testing.T) {
	assert.NotEmpty(t, useragent.NewUserAgentRandom().Call().UserAgent)
}
