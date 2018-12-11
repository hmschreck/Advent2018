package main

import (
	"fmt"
	"time"
)

const PLAYERS = 465
const HIGHEST = 71940 * 10

type Player struct {
	Score int
}

// The Marble type contains a value, as well as pointers to both its next and previous marbles.
// Importantly, the first marble has both a next and previous of itself.
type Marble struct {
	Value int
	Next  *Marble
	Prev  *Marble
}

// Takes the proper turn for the given value
func (marble *Marble) TakeTurn(value int) (score int, currentMarble *Marble) {
	// If the value is not divisible by 23, insert a new marble between
	// marble.Next and Marble.Next.Next (the next, and the one after that),
	// then fix pointers so that they point to the correct locations
	if value%23 != 0 {
		oneOver := marble.Next
		twoOver := marble.Next.Next
		currentMarble = &Marble{Value: value, Next: twoOver, Prev: oneOver}
		(*oneOver).Next = currentMarble
		(*twoOver).Prev = currentMarble
		return
		// If the value is divisibly by 23, get the score of the marble
		// 7 back (CCW), then set -8.Next to -6 and -6.Prev to -8.
	} else {
		removeMarble := marble.Prev.Prev.Prev.Prev.Prev.Prev.Prev
		score = removeMarble.Value + value
		removeMarble.Prev.Next = removeMarble.Next
		removeMarble.Next.Prev = removeMarble.Prev
		currentMarble = removeMarble.Next
		return
	}
}

func main() {
	// Create the the zero-value marble.
	currentMarble := &Marble{Value: 0}
	currentMarble.Next = currentMarble
	currentMarble.Prev = currentMarble
	players := make([]Player, PLAYERS)
	currentPlayer := 0
	// Evaluate turns

	timeStart := time.Now().UnixNano()
	// Loop through until we hit the highest value marble.
	for i := 1; i <= HIGHEST; i++ {
		if currentPlayer == PLAYERS {
			currentPlayer = 0
		}
		score, newCurrent := currentMarble.TakeTurn(i)
		players[currentPlayer].Score += score
		currentMarble = newCurrent
		currentPlayer += 1
	}
	max := 0
	for i := range players {
		if players[i].Score > max {
			max = players[i].Score
		}
	}
	fmt.Println(max)
	fmt.Println(float64(time.Now().UnixNano()-timeStart) / 1000000.0)
}
