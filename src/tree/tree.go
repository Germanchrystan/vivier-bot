package tree

//====================================================================//

type Tree struct {
	root            *Node
	scale           int
	indexLimitLeft  int
	indexLimitRight int
}

func NewTree(rootIndex, scale, indexLimitLeft, indexLimitRight int) *Tree {
	tree := &Tree{
		scale:           scale,
		root:            newNode(rootIndex),
		indexLimitLeft:  indexLimitLeft,
		indexLimitRight: indexLimitRight,
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
