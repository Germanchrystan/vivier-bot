package main

import (
	"github.com/Germanchrystan/vivier-forest/src/composer"
	"github.com/Germanchrystan/vivier-forest/src/writer"
)

func main() {
	/* tree := tree.NewTree(7, 2, 2, 2)
	fmt.Println(tree.GetTraversedTree())

	fmt.Println(prog)
	*/
	prog := composer.CreateChordSet(64, 5)
	baseProgStr, upperProgStr := writer.EngraveProgression(prog)
	parts := writer.EngraveParts(baseProgStr, upperProgStr)
	file := writer.EngraveFile(parts)
	writer.CreateFile(file)

}
