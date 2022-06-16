package main

import "fmt"

// Super basic tutorial (links to the part i'm on):
// https://kyxey.medium.com/lets-go-part-5-functions-6db1c36ff3bb

// "Level 2" CLI tutorial:
// https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f

func main() {
	fmt.Print("\n╠══╣ Simple Go Chess CLI ╠══╣\n\n")

  fmt.Print("It takes in a FEN, and prints it to the screen 😎\n")
  printFEN("Nn2k1nr/1p1q1ppp/8/2p4P/pPPp2b1/b2B1N2/P2P1PP1/R1BQ1RK1 b - b3 0 13")
}