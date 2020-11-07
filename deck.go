package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var d deck
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}
	for _, s := range cardSuits {
		for _, v := range cardValues {
			d = append(d, v+" of "+s)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toDeck(s string) deck {
	return strings.Split(s, ",")
}

func (d deck) saveDeck(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeck(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// my shuffle
func (d deck) shuffle(times int) {
	for n := 0; n < times; n++ {
		for i := range d {
			rand.Seed(time.Now().UnixNano())
			swap := rand.Intn(len(d) - 1)
			d[i], d[swap] = d[swap], d[i]
		}
	}
}

// course shuffle
func (d deck) shuffle2(times int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for n := 0; n < times; n++ {
		for i := range d {
			swap := r.Intn(len(d) - 1)
			d[i], d[swap] = d[swap], d[i]
		}
	}
}
