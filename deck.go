/**
* This is a deck.
* My first go thing so thing will change.
* I do not understand magit
**/
package deck

import (
	"errors"
	"fmt"
	"math/rand"
	slice "rezder.com/slice/int"
	"time"
)

type Deck struct {
	Perm     []int
	Delt     []bool
	Top      int
	Returned []int
}

func New(i int) *Deck {
	deck := Deck{
		Delt: make([]bool, i), // should all be zero value
		Top:  0,
	}
	rand.Seed(time.Now().UnixNano())
	deck.Perm = rand.Perm(i)
	return &deck
}
func (deck *Deck) Equal(other *Deck) (equal bool) {
	if other == nil && deck == nil {
		equal = true
	} else if other != nil && deck != nil {
		if deck == other {
			equal = true
		} else if other.Top == deck.Top {
			if slice.Equal(other.Perm, deck.Perm) && slice.Equal(other.Returned, deck.Returned) {
				equal = true
				for i, v := range other.Delt {
					if v != deck.Delt[i] {
						equal = false
						break
					}
				}
			}
		}
	}
	return equal
}
func (deck *Deck) Copy() (c *Deck) {
	if deck != nil {
		c = new(Deck)
		c.Top = deck.Top
		c.Perm = make([]int, len(deck.Perm))
		copy(c.Perm, deck.Perm)
		c.Delt = make([]bool, len(deck.Delt))
		copy(c.Delt, deck.Delt)
		c.Returned = make([]int, len(deck.Returned))
		copy(c.Returned, deck.Returned)
	}
	return c
}
func (deck *Deck) String() (txt string) {
	if deck != nil {
		last := "nil"
		if deck.Top != 0 {
			last = fmt.Sprint(deck.Perm[deck.Top-1])
		}
		txt = fmt.Sprintf("Deck{Cards %v, Index: %v Last %v}", len(deck.Perm), deck.Top, last)
	} else {
		txt = "Deck{nil}"
	}
	return txt
}

//Return the returned cards will added to the top and acces in reverse order,
func (deck *Deck) Return(ret []int) {
	if len(ret) != 0 {
		deck.Returned = ret
	}
}

func (deck *Deck) Deal() (ix int, err error) {
	if len(deck.Returned) != 0 {
		ix = deck.Returned[len(deck.Returned)-1]
		deck.Returned = deck.Returned[:len(deck.Returned)-1]
	} else {
		if deck.Top < len(deck.Perm) {
			ix = deck.Perm[deck.Top]
			deck.Delt[deck.Top] = true
			deck.Top = deck.Top + 1
		} else {
			ix = -1
			err = errors.New("Deck is empty")
		}
	}
	return ix, err
}
func (deck *Deck) DealCard(c int) (err error) {
	found := false
	if len(deck.Returned) != 0 {
		for i, v := range deck.Returned {
			if v == c {
				copy(deck.Returned[:i], deck.Returned[i+1:])
				found = true
			}
		}
	}
	if !found {
		var ix = -1
		for i, v := range deck.Perm {
			if v == c {
				ix = i
				break
			}
		}
		if ix != -1 {
			if deck.Delt[ix] {
				err = errors.New("Card is allready delt")
			} else {
				deck.Delt[ix] = true
			}
			deck.Delt[ix] = true
		} else {
			err = errors.New("Card do not exist")
		}
	}

	return err
}
func (deck *Deck) Scuffle() {
	rand.Seed(time.Now().UnixNano())
	deck.Perm = rand.Perm(len(deck.Delt))
	for i := range deck.Delt {
		deck.Delt[i] = false
	}
	deck.Top = 0
	deck.Returned = nil
}
func (deck *Deck) Empty() (empty bool) {
	if len(deck.Returned) == 0 {
		empty = true
		for _, delt := range deck.Delt {
			if !delt {
				empty = false
				break
			}
		}
	}
	return empty
}

//Remaining returns the no delt cards in the order they would be delt in.
func (deck *Deck) Remaining() (res []int) {
	res = make([]int, 0, len(deck.Perm)-deck.Top+len(deck.Returned))
	if len(deck.Returned) != 0 {
		for _, v := range deck.Returned {
			res = append(res, v)
		}
	}
	for i := deck.Top; i < len(deck.Perm); i++ {
		if !deck.Delt[i] {
			res = append(res, deck.Perm[i])
		}
	}

	return res
}
