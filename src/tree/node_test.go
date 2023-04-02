package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_neeNode(t *testing.T) {
	node := NewNode(1)
	assert.Equal(t, 1, node.index)
}

func Test_Populate1(t *testing.T) {
	node := NewNode(2)
	node.Populate(1, 1, 1)
	assert.Equal(t, 1, node.left.index)
	assert.Nil(t, node.right)
}

func Test_Populate2(t *testing.T) {
	node := NewNode(3)
	node.Populate(1, 1, 1)
	assert.Equal(t, 2, node.left.index)
	assert.Equal(t, 1, node.right.index)
}

func Test_Populate3(t *testing.T) {
	node := NewNode(4)
	node.Populate(2, 1, 1)
	assert.Equal(t, 2, node.left.index)
	assert.Nil(t, node.right)
}

func Test_Populate4(t *testing.T) {
	node := NewNode(4)
	node.Populate(2, 3, 1)
	assert.Nil(t, node.left)
}
