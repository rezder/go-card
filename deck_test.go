package deck

import "testing"

type card struct {
	color int
	value int
}

//TestDeck just a quick run throw
func TestDeck(t *testing.T) {
	t.Logf("Deck")
	deck := New(52)
	deck.Scuffle()
	d52 := createPlayDeck()
	logDeck(deck, d52, t)
	t.Logf("Scuffle")
	deck.Scuffle()
	logDeck(deck, d52, t)

}
func logDeck(d *Deck, deck52 []card, t *testing.T) {
	for i := 0; i < 52; i++ {
		ix, _ := d.Deal()
		t.Logf("Card %v, %v", co[deck52[ix].color], deck52[ix].value)
	}
}
func createPlayDeck() (d []card) {
	d = make([]card, 52)
	for c := 0; c < 4; c++ {
		for i := 2; i < 15; i++ {
			d[c*13+i-2] = card{c, i}
		}
	}
	return d
}

var co = [4]string{
	"\u2663",
	"\u2666",
	"\u2665",
	"\u2660",
}
