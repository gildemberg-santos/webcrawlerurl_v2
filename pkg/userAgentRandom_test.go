package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAgentRandom_Call(t *testing.T) {
	userAgentRandom := UserAgentRandom{}
	userAgentRandom.Call()
	assert.NotEmpty(t, userAgentRandom.UserAgent)
}
