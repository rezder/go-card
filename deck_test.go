package deck

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"
)

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
	deck.Scuffle()
	ix, _ := deck.Deal()
	t.Logf("Deal card %v,%v\n", co[d52[ix].color], d52[ix].value)
	r := make([]int, 2)
	r = append(r, ix)
	ix, _ = deck.Deal()
	t.Logf("Deal card %v,%v\n", co[d52[ix].color], d52[ix].value)
	r = append(r, ix)
	deck.Return(r)
	logDeck(deck, d52, t)
	t.Logf("Returned\n")
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

func TestSave(t *testing.T) {
	fileName := "test/save.gob"
	deck := New(10)
	file, err := os.Create(fileName)
	//f,err:=os.Open(fileName)

	if err == nil {
		err = Save(deck, file)
		file.Close()
		if err != nil {
			t.Errorf("Save game file error. File :%v, Error: %v", fileName, err.Error())
		} else {
			file, err = os.Open(fileName)
			saveDeck, err := Load(file)
			if err != nil {
				t.Errorf("Load game file error. File :%v, Error: %v", fileName, err.Error())
			} else {
				if !saveDeck.Equal(deck) {
					t.Error("Save and load game not equal")
				}
			}
		}
	} else {
		t.Errorf("Create file error. File :%v, Error: %v", fileName, err.Error())
	}
}
func Save(game *Deck, file *os.File) (err error) {
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(game)

	return err
}
func Load(file *os.File) (deck *Deck, err error) {
	decoder := gob.NewDecoder(file)
	var d Deck
	err = decoder.Decode(&d)
	deck = &d
	return deck, err
}
func TestDecoder(t *testing.T) {
	deck := New(10)
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(deck)
	if err != nil {
		t.Errorf("Error encoding")
	}

	var loadDeck Deck
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&loadDeck)
	if err != nil {
		t.Errorf("Error decoding")
	} else {
		if !deck.Equal(&loadDeck) {
			t.Logf("Deck :%v\nLoad :%v", deck, loadDeck)
			t.Error("Save and load deck not equal")
		}
	}
}
func TestFormate(t *testing.T) {
	value := *New(10)
	value.Deal()
	t.Logf("Value: %v,Pointer: %v", value, &value)
	//	t.Error(" Fail")
}
