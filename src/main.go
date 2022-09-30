package main

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/tree"
)

func main() {
	tree := tree.NewTree(7, 2, 2, 2)
	fmt.Println(tree.GetTraversedTree())
}
