package schema

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)

	user := &User{
		ID:        "test",
		Pass:      "Pass",
		IsBlocked: false,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	str := user.String()
	assert.True(strings.Contains(str, `"id":"test"`))
	assert.True(strings.Contains(str, `"pass":"Pass"`))
	assert.True(strings.Contains(str, `"isBlocked":false`))
	assert.True(strings.Contains(str, `"created":"20`))
	assert.True(strings.Contains(str, `"updated":"20`))

	str = user.Result().String()
	assert.True(strings.Contains(str, `"id":"test"`))
	assert.False(strings.Contains(str, `"pass":"Pass"`))
	assert.False(strings.Contains(str, `"isBlocked":false`))
	assert.True(strings.Contains(str, `"created":"20`))
	assert.True(strings.Contains(str, `"updated":"20`))
}
