package rawDelList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawDelList(t *testing.T) {
	intList := []int{3, 2, 3, 1, 1, 2}
	assert.Equal(t, RawDelList(intList, 2), 4)
}
