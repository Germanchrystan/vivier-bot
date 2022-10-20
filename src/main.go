package main

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/composer"
)

func main() {
	/* tree := tree.NewTree(7, 2, 2, 2)
	// fmt.Println(tree.GetTraversedTree())

	coords := composer.CreateHLBaseNotes()
	fmt.Println(coords)
	*/

	prog := composer.CreateChordSet(3, 3)
	fmt.Println(prog)
}
