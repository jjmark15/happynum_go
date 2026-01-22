package happynum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFirstIteration(t *testing.T) {
	assert.True(t, isFirstIteration(1234))
	assert.True(t, isFirstIteration(0000))
	assert.True(t, isFirstIteration(123))
	assert.False(t, isFirstIteration(1230))
	assert.False(t, isFirstIteration(1243))
}

func TestIsUnhappy(t *testing.T) {
	assert.False(t, IsHappy(89))
}

func TestIsHappy(t *testing.T) {
	assert.True(t, IsHappy(10))
}

func TestSquareSum(t *testing.T) {
	assert.Equal(t, 1, squareSum(1))
	assert.Equal(t, 4, squareSum(2))
	assert.Equal(t, 20, squareSum(204))
	assert.Equal(t, 0, squareSum(0))
}

func TestDistinctHappyRangeCount(t *testing.T) {
	assert.Equal(t, 711, DistinctHappyRangeCount(1, 1000000))
}
