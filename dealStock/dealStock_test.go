package dealStock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDealStock(t *testing.T) {
	prices := []int32{7, 2, 1, 8, 3, 9}
	assert.Equal(t, DealStock(prices), int32(13), "they should equal")
	assert.Equal(t, DealStock1(prices), int32(13), "they should equal")
}
