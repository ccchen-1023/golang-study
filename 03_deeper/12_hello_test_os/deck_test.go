package main

import (
	"os"
	"testing"
)

// TIP: Test commands:
// Test file name must be end of _test.go
// run package tests: go test
// run file tests: go test deck_test.go deck.go
// run a test function: go test -run Test_NewDeck_Size

// Test func name must be start with Test
func Test_newDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expect deck size is 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect first card is Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expect first card is K of Clubs, but got %v", d[len(d)-1])
	}
}

// TIP: No meanings to use multiple func for testing number. Use -cover instead
// Coverage commands:
// test coverage: go test -cover
// output test coverage: go test -coverprofile=size_coverage.out
// functions coverage from output file: go tool cover -func=size_coverage.out
// functions coverage(html) from output file: go tool cover -html=size_coverage.out
/*
func Test_NewDeck_First(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expect first card is Ace of Spades, but got %v", d[0])
	}
}

func Test_NewDeck_Last(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expect first card is K of Clubs, but got %v", d[len(d)-1])
	}
}
*/

func Test_saveToFile_newDeckFromFile(t *testing.T) {
	fileName := "temp_test_file"
	os.Remove(fileName)
	d := newDeck()
	d.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)

	if d.toString() != loadedDeck.toString() {
		t.Errorf("Expect file contents of newDeck() and newDeck() are same, but diff.")
	}
	os.Remove(fileName)
}

// TIP: Benchmark command:
// run a function benchmark: go test -bench Benchmark_shuffle
// run package benchmarks: go test -bench .
// see: https://openhome.cc/Gossip/Go/Testing.html

func Benchmark_shuffle(b *testing.B) {
	d := newDeck()
	for i := 0; i < b.N; i++ {
		d.shuffle()
	}
}

func Benchmark_toString(b *testing.B) {
	d := newDeck()
	for i := 0; i < b.N; i++ {
		d.toString()
	}
}
