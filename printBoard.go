package main

import (
  "fmt"
  "regexp"
  "strings"
)

// how to erase stuff from the console... https://stackoverflow.com/questions/48190115/golang-erase-symbols-from-console

// regex checking the fen validity: \s*([rnbqkpRNBQKP1-8]+\/){7}([rnbqkpRNBQKP1-8]+)\s[bw-]\s(([a-hkqA-HKQ]{1,4})|(-))\s(([a-h][36])|(-))\s\d+\s\d+\s*

var pieces = map[string]string {
	"K": "♔",
	"Q": "♕",
	"R": "♖",
	"B": "♗",
	"N": "♘",
	"P": "♙",
	"k": "♚",
	"q": "♛",
	"r": "♜",
	"b": "♝",
	"n": "♞",
	"p": "♟",
}

// Takes a FEN string; returns whether it's formed properly
func validateFEN(boardState string) bool {
  re := regexp.MustCompile(`\s*([rnbqkpRNBQKP1-8]+\/){7}([rnbqkpRNBQKP1-8]+)\s[bw-]\s(([a-hkqA-HKQ]{1,4})|(-))\s(([a-h][36])|(-))\s\d+\s\d+\s*`)
  return re.MatchString(boardState)
}
 
// Takes a FEN string; prints the board accordingly
func printFEN(boardState string) {
  if (!validateFEN(boardState)) {
    return
  }

  values := strings.Split(boardState, " ")
	fmt.Printf("\nPrinting board from: %v", values[3])
  fmt.Printf(pieces["K"])

}

// Takes a PGN string; prints the board accordingly
func printPGN(boardPGN string) {

}