package tree

//====================================================================//

type Tree struct {
	root            *Node
	scale           int
	indexLimitLeft  int
	indexLimitRight int
	left            *Tree
	right           *Tree
}

func NewTree(rootIndex, scale, indexLimitLeft, indexLimitRight int) *Tree {
	tree := &Tree{
		scale:           scale,
		root:            NewNode(rootIndex),
		indexLimitLeft:  indexLimitLeft,
		indexLimitRight: indexLimitRight,
		left:            nil,
		right:           nil,
	}

	mayKeepTraversing := true
	for mayKeepTraversing {
		mayKeepTraversing = tree.root.Populate(scale, indexLimitLeft, indexLimitRight)
	}

	return tree
}

func (t *Tree) GetTraversedTree() []int {
	return t.root.inOrderTree()
}

//====================================================================//
