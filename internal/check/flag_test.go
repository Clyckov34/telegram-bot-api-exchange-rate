package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	var data string = "Test"

	err := Flag(data)
	assert.Nil(t, err)
}
