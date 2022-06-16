package main

import (
  "fmt"
  "regexp"
  "strings"
)

// how to erase stuff from the console... https://stackoverflow.com/questions/48190115/golang-erase-symbols-from-console

// regex checking the fen validity: \s*([rnbqkpRNBQKP1-8]+\/){7}([rnbqkpRNBQKP1-8]+)\s[bw-]\s(([a-hkqA-HKQ]{1,4})|(-))\s(([a-h][36])|(-))\s\d+\s\d+\s*

const asciiPieces = true

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

  // [board, toMove, castling, epSquare, halfmoveClock, moveNum]
  fen := strings.Split(boardState, " ")
  var extraStuff = [2]string {"", ""}
  if asciiPieces {
    extraStuff[0] = " "
    extraStuff[1] = "═"
  }

  var rank = 8;
  var turnMessage = "White to move."
  if (fen[1] == "b") {
    turnMessage = "Black to move."
  }
  var castlingPermissions = "Castling Permissions: " + fen[2]
  var rankMessages = [8]string {"", "", "", "", "", "", "", turnMessage}

  fmt.Printf("\n  ╔%v════════════════════════╗\n8 ║", extraStuff[1])
  for _, char := range fen[0] {
    if (char == '/') {
      rank--
      fmt.Printf("%v║  %v\n%v ║", extraStuff[0], rankMessages[rank], rank)
    } else if piece, isPiece := pieces[string(char)]; isPiece {
      if (asciiPieces) {
        fmt.Print(" ", piece, " ")
      } else {
        fmt.Print(" ", string(char), " ")
      }
    } else {
      for i:= 0; i < int(char - '0'); i++ {
        fmt.Print(" · ")
      }
    }
  }
  fmt.Printf("%v║  %v\n  ╚%v════════════════════════╝\n    a  b  c  d  e  f  g  h\n", extraStuff[0], castlingPermissions, extraStuff[1])

	fmt.Print("\nPrinting board from this FEN:\n", boardState, "\n\n")

}