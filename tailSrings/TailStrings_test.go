package tailStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTailStrings(t *testing.T) {
	str := "hello  wuhaurou "
	assert.Equal(t, TailStrings(str), 8)
}
