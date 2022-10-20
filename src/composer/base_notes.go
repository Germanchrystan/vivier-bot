package composer

import (
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/wesovilabs/koazee"
)

type NoteDirection int

const (
	up   NoteDirection = 1
	down               = -1
)

type Interval int

const (
	minorSecond     Interval = 1
	mayorSecond              = 2
	minorThird               = 3
	majorThird               = 4
	perfectFourth            = 5
	augmentedFourth          = 6
)

type baseNoteMovement struct {
	Direction NoteDirection
	Interval  Interval
}

func createIntervalHierarchy() []wr.Choice {
	var intervalNumbers []int
	var baseNoteMovements []baseNoteMovement
	stream := koazee.StreamOf(intervalNumbers)
	for len(intervalNumbers) < 3 {
		newInt := getRandomIntBetweenRanges(1, 6)
		contains, _ := stream.Contains(newInt)
		if !contains {
			intervalNumbers = append(intervalNumbers, newInt)

			baseNoteMovements = append(
				baseNoteMovements,
				baseNoteMovement{Interval: Interval(newInt), Direction: down},
				baseNoteMovement{Interval: Interval(newInt), Direction: up},
			)
			stream = koazee.StreamOf(intervalNumbers)
		}
	}
	return CreateChoices(baseNoteMovements)
}

func isSameNote(note1, note2 Note) bool {
	if note1.Step == note2.Step && note1.Accidental == note2.Accidental {
		return true
	}
	return false
}

func CreateHLBaseNotes() [2]Note {
	var result [2]Note
	bassNote := getRandomNoteBetweenRange(0, 7, BassNoteCatalog)
	celloNote := getRandomNoteBetweenRange(len(CelloNoteCatalog)-8, len(CelloNoteCatalog)-1, CelloNoteCatalog)
	result[0] = bassNote
	result[1] = celloNote
	return result
}

func createFollowingBaseNotes(prevBaseNotes [2]Note, chooser *wr.Chooser) [2]Note {
	rand.Seed(time.Now().UTC().UnixNano())

	var result [2]Note
	noteMovement := chooser.Pick().(baseNoteMovement)                    // Pick base note movement
	interval := int(noteMovement.Interval) * int(noteMovement.Direction) // get int from baseNoteMovement struct

	randomIndex := rand.Intn(2) // Random pick between taking bass or cello first (50/50)

	if randomIndex == 0 {
		isValid, newNoteIndex := isValidNoteMovement(BassNoteCatalog, prevBaseNotes[0], interval)
		if isValid && !isSameNote(BassNoteCatalog[newNoteIndex], prevBaseNotes[1]) {
			return [2]Note{BassNoteCatalog[newNoteIndex], prevBaseNotes[1]}
		} else if isValid, newNoteIndex = isValidNoteMovement(CelloNoteCatalog, prevBaseNotes[1], interval); isValid && !isSameNote(prevBaseNotes[0], CelloNoteCatalog[newNoteIndex]) {
			return [2]Note{prevBaseNotes[0], CelloNoteCatalog[newNoteIndex]}
		} else {
			return createFollowingBaseNotes(prevBaseNotes, chooser)
		}
	} else {
		isValid, newNoteIndex := isValidNoteMovement(CelloNoteCatalog, prevBaseNotes[1], interval)
		if isValid && !isSameNote(prevBaseNotes[0], CelloNoteCatalog[newNoteIndex]) {
			return [2]Note{prevBaseNotes[0], CelloNoteCatalog[newNoteIndex]}
		} else if isValid, newNoteIndex = isValidNoteMovement(BassNoteCatalog, prevBaseNotes[0], interval); isValid && !isSameNote(BassNoteCatalog[newNoteIndex], prevBaseNotes[1]) {
			return [2]Note{BassNoteCatalog[newNoteIndex], prevBaseNotes[1]}
		} else {
			return createFollowingBaseNotes(prevBaseNotes, chooser)
		}
	}
	return result
}

func isValidNoteMovement(noteCatalog []Note, note Note, newInterval int) (bool, int) {
	stream := koazee.StreamOf(noteCatalog)
	index, _ := stream.IndexOf(note)

	if index < 0 {
		return false, index
	}

	newIntervalIndex := index + newInterval
	if newIntervalIndex >= 0 && newIntervalIndex <= len(noteCatalog)-1 {
		return true, newIntervalIndex
	}

	return false, newIntervalIndex
}

func getRandomNoteBetweenRange(min, max int, catalog []Note) Note {
	index := getRandomIntBetweenRanges(min, max)
	return catalog[index]
}

func CreateBaseNoteProgression(progressionLen int) [][2]Note {
	var result [][2]Note
	intervalHierarchy := createIntervalHierarchy()
	chooser, err := CreateNewChooser(intervalHierarchy)
	if err != nil {
		panic(err)
	}
	firstBaseNotePair := CreateHLBaseNotes()
	result = append(result, firstBaseNotePair)

	for len(result) < progressionLen {
		prevPair := result[len(result)-1]
		newPair := createFollowingBaseNotes(prevPair, chooser)
		result = append(result, newPair)
	}
	return result
}
