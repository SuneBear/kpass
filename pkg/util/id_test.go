package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOID(t *testing.T) {
	assert := assert.New(t)

	id := NewOID()
	id2 := id

	assert.True(id.Equal(id2))
	assert.False(id.Equal(NewOID()))
	assert.Equal(time.Now().Unix(), id.Time().Unix())

	i, err := ParseOID("")
	assert.NotNil(err)
	assert.False(i.Valid())

	i, err = ParseOID("abc")
	assert.NotNil(err)
	assert.False(i.Valid())

	i, err = ParseOID("1234567890")
	assert.NotNil(err)
	assert.False(i.Valid())

	s, _ := (&id).MarshalText()
	i, err = ParseOID(string(s))
	assert.True(i.Valid())
}
