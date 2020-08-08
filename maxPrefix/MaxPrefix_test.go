package maxPrefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPrefix(t *testing.T) {
	abc := []string{"hello", "hello world", "he"}
	assert.Equal(t, MaxPrefix(abc), "he", "they should equal")
}
