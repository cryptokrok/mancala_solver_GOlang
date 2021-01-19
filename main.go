package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kNoAPP/mancala/pkg/mancala_solver"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader
var sample mancala_solver.MancalaState
var depth uint8 = 6

func main() {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("\n\nWelcome to Mancala's Adversarial Lookup")
	fmt.Println("  #(0-13)   : Indicate player choice")
	fmt.Println("  r         : Reset")
	fmt.Println("----------------------------------------")


	for {
		coreGameLoop()
		fmt.Print("\n\nResetting automagically...\n")
	}
}

func coreGameLoop() {
	sample.Board = [14]uint8{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
	// Keep trying until correct input is achieved
	for err := setDepth(); err != nil; err = setDepth() {}
	for err := getStarter(); err != nil; err = getStarter() {}

	var ally float32
	var adversary float32
	var end bool
	for end, ally, adversary = sample.IsEndOfGame(); !end; end, ally, adversary = sample.IsEndOfGame() {
		printGameStatus()
		var reset bool
		var err error
		for reset, err = getMove(); err != nil; reset, err = getMove() {}
		if reset {
			return
		}
		fmt.Println("Move made! Updating the game...")
	}

	fmt.Println("Game Completed! Time for the Final Scores.")
	fmt.Println("You: ", ally)
	fmt.Println("Them: ", adversary)
}

func getMove() (reset bool, err error) {
	fmt.Print("Type a number representing the move (0-12) or 'r' to reset: ")
	in, _ := reader.ReadString('\n')
	in = strings.TrimSpace(in)

	if strings.EqualFold("r", in) {
		return true, nil
	} else if s, err := strconv.ParseUint(in, 10, 8); err == nil && s <= 12 {
		sample, err = mancala_solver.AdvanceState(sample, uint8(s))
		if err != nil {
			return false, errors.New("invalid turn")
		}
		return false, nil
	}

	return false, errors.New("invalid input")
}

func printGameStatus() {
	alliedScore, adversaryScore, moves := mancala_solver.CalculateMove(sample, depth)
	fmt.Print("\nBoard\n")
	sample.PrintBoard()
	fmt.Println("----------------------------")
	fmt.Println("Your Turn?        : ", sample.AlliedTurn)
	fmt.Println("Allied Score      : ", alliedScore)
	fmt.Println("Adversarial Score : ", adversaryScore)
	fmt.Print("Best Moves        :  ")
	for e := moves.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Print("\n")
}

func getStarter() error {
	fmt.Print("Are you going first? (y/n): ")
	in, _ := reader.ReadString('\n')
	in = strings.TrimSpace(in)

	if strings.EqualFold("y", in) {
		sample.AlliedTurn = true
		return nil
	} else if strings.EqualFold("n", in) {
		sample.AlliedTurn = false
		return nil
	}

	return errors.New("invalid input")
}

func setDepth() error {
	fmt.Print("Please set a search depth. The higher, the smarter, the slower. (rec. 10): ")
	in, _ := reader.ReadString('\n')
	in = strings.TrimSpace(in)

	if s, err := strconv.ParseUint(in, 10, 8); err == nil {
		depth = uint8(s)
		return nil
	}

	return errors.New("invalid input")
}