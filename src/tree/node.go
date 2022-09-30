package tree

type Node struct {
	index int
	value int
	right *Node
	left  *Node
}

func newNode(index int) *Node {
	return &Node{
		index: index,
		value: GetFibValueByIndex(index),
		left:  nil,
		right: nil,
	}
}

func (n *Node) Populate(scale, indexLimitLeft, indexLimitRight int) bool {
	mayKeepTraversing := false
	if n.left == nil {
		if n.index-scale >= indexLimitLeft { // If the index of the node minus the scale is bigger than the limit
			n.left = newNode(n.index - scale)
			mayKeepTraversing = true
		} else {
			n.left = nil
		}
	} else {
		n.left.Populate(scale, indexLimitLeft, indexLimitRight) // If Left node is not empty, use recursion
	}

	if n.right == nil {
		if n.index-(scale*2) >= indexLimitRight {
			n.right = newNode(n.index - (scale * 2))
			mayKeepTraversing = true
		} else {
			n.right = nil
		}
	} else {
		n.right.Populate(scale, indexLimitLeft, indexLimitRight)
	}
	return mayKeepTraversing
}

func (n *Node) inOrderTree() []int {
	var arr []int
	if n.left != nil {
		arr = append(arr, n.left.inOrderTree()...)
	}
	arr = append(arr, n.value)
	if n.right != nil {
		arr = append(arr, n.right.inOrderTree()...)
	}
	return arr
}
