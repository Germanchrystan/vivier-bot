package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTree(t *testing.T) {
	tree := NewTree(2, 1, 1, 1)
	assert.Equal(t, 1, tree.root.left.index)
	assert.Nil(t, tree.root.right)
}

func Test_GetTraversedTree1(t *testing.T) {
	tree := NewTree(2, 1, 1, 1)
	want := []int{1, 2}
	got := tree.GetTraversedTree()
	assert.Equal(t, want, got)
}

func Test_GetTraversedTree2(t *testing.T) {
	tree := NewTree(3, 1, 1, 1)
	want := []int{1, 2, 3, 1}
	got := tree.GetTraversedTree()
	assert.Equal(t, want, got)
}

func Test_GetTraversedTree3(t *testing.T) {
	tree := NewTree(4, 1, 1, 1)
	want := []int{1, 2, 3, 1, 4, 1, 2}
	got := tree.GetTraversedTree()
	assert.Equal(t, want, got)
}

func Test_GetTraversedTree4(t *testing.T) {
	tree := NewTree(5, 1, 2, 2)
	want := []int{2, 3, 4, 2, 5, 2, 3}
	got := tree.GetTraversedTree()
	assert.Equal(t, want, got)
}

func Test_GetTraversedTree5(t *testing.T) {
	tree := NewTree(10, 2, 2, 2)
	want := []int{2, 4, 6, 2, 8, 2, 4, 10, 2, 4, 6, 2}
	got := tree.GetTraversedTree()
	assert.Equal(t, want, got)
}
