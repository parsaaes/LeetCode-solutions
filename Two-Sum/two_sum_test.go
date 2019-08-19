package Two_Sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	result := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(result, []int{0, 1})
}
