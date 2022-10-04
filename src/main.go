package main

import (
	"fmt"

	"github.com/Germanchrystan/vivier-forest/src/writer"
)

func main() {
	// tree := tree.NewTree(7, 2, 2, 2)
	// fmt.Println(tree.GetTraversedTree())

	noteFound := writer.FindNearestNote(2000, writer.NotesCatalog)
	fmt.Println(noteFound)
}
