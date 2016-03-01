/**
* This is a deck.
* My first go thing so thing will change.
* I do not understand magit
**/
package deck

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	perm []int
	delt []bool
	top  int
}

func New(i int) *Deck {
	deck := Deck{
		perm: rand.Perm(i),
		delt: make([]bool, i), // should all be zero value
		top:  0,
	}
	return &deck
}
func (deck *Deck) Deal() (ix int, err error) {
	if deck.top < len(deck.perm) {
		ix = deck.perm[deck.top]
		deck.delt[deck.top] = true
		deck.top = deck.top + 1
	} else {
		ix = -1
		err = errors.New("Deck is empty")
	}
	return ix, err
}
func (deck *Deck) DealCard(c int) (err error) {
	var ix = -1
	for i, v := range deck.perm {
		if v == c {
			ix = i
			break
		}
	}
	if ix != -1 {
		if deck.delt[ix] {
			err = errors.New("Card is allready delt")
		} else {
			deck.delt[ix] = true
		}
		deck.delt[ix] = true
	} else {
		err = errors.New("Card do not exist")
	}
	return err
}
func (deck *Deck) Scuffle() {
	rand.Seed(time.Now().UnixNano())
	deck.perm = rand.Perm(len(deck.delt))
	for i := range deck.delt {
		deck.delt[i] = false
	}
	deck.top = 0
}
