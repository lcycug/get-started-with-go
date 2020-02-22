package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
<<<<<<< HEAD
	"time"
=======
>>>>>>> d1f9b1c... functionalities of deck
)

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
<<<<<<< HEAD
			cards = append(cards, value+" of "+suite)
=======
			cards = append(cards, suite+" of "+value)
>>>>>>> d1f9b1c... functionalities of deck
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) save2File(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
<<<<<<< HEAD
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		position := r.Intn(len(d) - 1)
=======
	for index := range d {
		position := rand.Intn(len(d) - 1)
>>>>>>> d1f9b1c... functionalities of deck
		d[index], d[position] = d[position], d[index]
	}
}
