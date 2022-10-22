package composer

func CreateColorSet(colorSetLength int) (uniqueColor []Vector, colorPossibilities [][]Vector) {
	/*
		For now, HL Chord is the only one with
		a unique color, made with 9 vector points
		and 5 vectors
	*/
	uniqueColor = NewColor(9, 5)
	var colors [][]Vector
	for len(colors) < colorSetLength-1 {
		randomVectorPoints := getRandomIntBetweenRanges(3, 7)
		randomVectorQuantity := getRandomIntBetweenRanges(2, 4)
		newColor := NewColor(randomVectorPoints, randomVectorQuantity)
		colors = append(colors, newColor)
	}

	return uniqueColor, colors
}

func CreateChordSet(chordSetlength, colorSetLength int) []Chord {
	baseNotes := CreateBaseNoteProgression(chordSetlength)

	uniqueColor, colors := CreateColorSet(colorSetLength)
	choices := CreateChoices[[]Vector](colors)
	chooser, err := CreateNewChooser(choices)
	if err != nil {
		panic(err)
	}

	var result []Chord
	for index, bn := range baseNotes {
		var color []Vector
		if index == 0 {
			color = uniqueColor
		} else {
			color = chooser.Pick().([]Vector)
		}
		newChord := NewChord(color, bn)
		result = append(result, newChord)
	}
	return result
}
