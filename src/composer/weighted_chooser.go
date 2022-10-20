package composer

import wr "github.com/mroth/weightedrand"

func CreateChoices[T baseNoteMovement | []Vector](possibilities []T) []wr.Choice {
	var result []wr.Choice
	choicePoints := 100
	for _, interval := range possibilities {
		var points int

		if choicePoints/2 <= 10 {
			points = choicePoints
		} else {
			points = getRandomIntBetweenRanges(10, choicePoints/2)

		}
		result = append(result, wr.Choice{Item: interval, Weight: uint(points)})
		choicePoints -= points
	}
	return result
}

func CreateNewChooser(choices []wr.Choice) (*wr.Chooser, error) {
	return wr.NewChooser(choices...)
}
