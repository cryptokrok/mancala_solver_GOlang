package mancala_solver

import (
	"container/list"
	"errors"
	"fmt"
)

type MancalaState struct {
	Board      [14]uint8
	AlliedTurn bool
}

func CalculateMove(state MancalaState, depth uint8) (alliedScore float32, adversaryScore float32, moves *list.List) {
	if end, allied, adversary := state.IsEndOfGame(); end {
		l := list.New()
		l.PushBack(-1)
		return allied, adversary, l
	}

	if depth == 0 {
		allied, adversary := state.evaluateBoard()
		return allied, adversary, list.New()
	}

	var base uint8
	if state.AlliedTurn {
		base = 0
	} else {
		base = 7
	}

	var bestMoves *list.List
	bestAllied, bestAdversary, bestMove := float32(0), float32(0), uint8(0)
	for i := uint8(0); i < 6; i++ {
		if state.Board[base+i] == 0 {
			continue
		}

		thisState, _ := AdvanceState(state, base+i)
		// Errors won't happen here since this is a controlled environment.

		thisAllied, thisAdversary, thisMoves := CalculateMove(thisState, depth-1)
		if state.AlliedTurn {
			if thisAllied > bestAllied {
				bestAllied = thisAllied
				bestAdversary = thisAdversary
				bestMoves = thisMoves
				bestMove = base + i
			}
		} else {
			if thisAdversary > bestAdversary {
				bestAllied = thisAllied
				bestAdversary = thisAdversary
				bestMoves = thisMoves
				bestMove = base + i
			}
		}
	}
	bestMoves.PushFront(bestMove)
	return bestAllied, bestAdversary, bestMoves
}

func AdvanceState(state MancalaState, moveIndex uint8) (nextState MancalaState, err error) {
	if (state.AlliedTurn && moveIndex >= 6) || (!state.AlliedTurn && (moveIndex <= 6 || moveIndex >= 13)) {
		return state, errors.New("invalid move")
	}

	if state.Board[moveIndex] == 0 {
		return state, errors.New("invalid move")
	}

	nextState = state
	nextState.AlliedTurn = !state.AlliedTurn
	nextState.Board[moveIndex] = 0
	currentPlacement := moveIndex
	for i := uint8(0); i < state.Board[moveIndex]; i++ {
		currentPlacement++
		if (state.AlliedTurn && currentPlacement == 13) || (!state.AlliedTurn && currentPlacement == 6) {
			currentPlacement++
		}
		currentPlacement %= 14

		nextState.Board[currentPlacement] = nextState.Board[currentPlacement] + 1
	}
	if state.AlliedTurn {
		if currentPlacement == 6 {
			nextState.AlliedTurn = state.AlliedTurn
		} else if isSlotAllied(currentPlacement) && nextState.Board[currentPlacement] == 1 && nextState.Board[getOppositeSlot(currentPlacement)] >= 1 {
			nextState.Board[6] += nextState.Board[getOppositeSlot(currentPlacement)] + 1
			nextState.Board[currentPlacement] = 0
			nextState.Board[getOppositeSlot(currentPlacement)] = 0
		}
	} else {
		if currentPlacement == 13 {
			nextState.AlliedTurn = state.AlliedTurn
		} else if !isSlotAllied(currentPlacement) && nextState.Board[currentPlacement] == 1 && nextState.Board[getOppositeSlot(currentPlacement)] >= 1 {
			nextState.Board[13] += nextState.Board[getOppositeSlot(currentPlacement)] + 1
			nextState.Board[currentPlacement] = 0
			nextState.Board[getOppositeSlot(currentPlacement)] = 0
		}
	}

	return nextState, nil
}

func (state *MancalaState) PrintBoard() {
	for i := 13; i > 6; i-- {
		fmt.Print(state.Board[i], " ")
	}
	fmt.Print("\n  ")
	for i := 0; i < 7; i++ {
		fmt.Print(state.Board[i], " ")
	}
	fmt.Print("\n")
}

func isSlotAllied(slot uint8) bool {
	return slot <= 6
}

func getOppositeSlot(slot uint8) uint8 {
	if slot != 6 && slot != 13 {
		return 12 - slot
	}

	return 0
}

func (state *MancalaState) evaluateBoard() (alliedScore float32, adversaryScore float32) {
	alliedScore = 0
	for i := 0; i < 6; i++ {
		alliedScore += float32(state.Board[i]) * 0.85
	}

	adversaryScore = 0
	for i := 7; i < 13; i++ {
		adversaryScore += float32(state.Board[i]) * 0.85
	}

	return alliedScore + float32(state.Board[6]), adversaryScore + float32(state.Board[13])
}

func (state *MancalaState) IsEndOfGame() (end bool, alliedScore float32, adversaryScore float32) {
	alliedScore = 0
	for i := 0; i < 6; i++ {
		alliedScore += float32(state.Board[i])
	}

	adversaryScore = 0
	for i := 7; i < 13; i++ {
		adversaryScore += float32(state.Board[i])
	}

	if alliedScore != 0 && adversaryScore != 0 {
		return false, alliedScore + float32(state.Board[6]), adversaryScore + float32(state.Board[13])
	}

	return true, alliedScore + float32(state.Board[6]), adversaryScore + float32(state.Board[13])
}
