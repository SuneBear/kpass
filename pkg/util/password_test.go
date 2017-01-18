package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandPass(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(10, len(RandPass(10)))
	assert.Equal(4, len(RandPass(2)))
	assert.Equal(64, len(RandPass(200)))

	fmt.Println(RandPass(12, 2, 2))
}
