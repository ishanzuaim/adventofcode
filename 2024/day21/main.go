package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var keypad [][]string = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"#", "0", "A"},
}

var dirs [][]string = [][]string{
	{"#", "^", "A"},
	{"<", "v", ">"},
}

type GameSettings struct {
	currPosMain Position
	currPosA    Position
	currPosB    Position
	output      []string
	total       []string
}

var keyMap map[string]string = map[string]string{
	"h": "<",
	"j": "v",
	"k": "^",
	"l": ">",
	"a": "A",
}

func readChar() string {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b[0])
}

func main() {
	gs := GameSettings{
		currPosMain: Position{2, 3},
		currPosA:    Position{2, 0},
		currPosB:    Position{2, 0},
		total:       make([]string, 0),
	}
	gs.print()
	// var move string
	for true {
		fmt.Print("Enter move: ")
		move := readChar()
		if move == "r" {
			gs = GameSettings{
				currPosMain: Position{2, 3},
				currPosA:    Position{2, 0},
				currPosB:    Position{2, 0},
				total:       make([]string, 0),
			}
      gs.print()
			continue
		} else if move == "q" {
			break
		}
		for _, val := range strings.Split(move, "") {
			gs.handleMove(keyMap[val])
		}
		gs.print()
	}
	fmt.Println(gs.total)
	fmt.Println(len(gs.total))
}

func (gs *GameSettings) handleMove(move string) {
	gs.total = append(gs.total, move)
	if move == ">" {
		gs.currPosB.x += 1
	}

	if move == "^" {
		gs.currPosB.y -= 1
	}

	if move == "v" {
		gs.currPosB.y += 1
	}

	if move == "<" {
		gs.currPosB.x -= 1
	}

	if move == "A" {
		gs.handleMoveA(dirs[gs.currPosB.y][gs.currPosB.x])
	}
}

func (gs *GameSettings) handleMoveA(move string) {
	if move == ">" {
		gs.currPosA.x += 1
	}

	if move == "^" {
		gs.currPosA.y -= 1
	}

	if move == "v" {
		gs.currPosA.y += 1
	}

	if move == "<" {
		gs.currPosA.x -= 1
	}

	if move == "A" {
		gs.handleMoveMain(dirs[gs.currPosA.y][gs.currPosA.x])
	}
}

func (gs *GameSettings) handleMoveMain(move string) {
	if move == ">" {
		gs.currPosMain.x += 1
	}

	if move == "^" {
		gs.currPosMain.y -= 1
	}

	if move == "v" {
		gs.currPosMain.y += 1
	}

	if move == "<" {
		gs.currPosMain.x -= 1
	}

	if move == "A" {
		gs.output = append(gs.output, keypad[gs.currPosMain.y][gs.currPosMain.x])
	}
}

func (gs GameSettings) print() {
	fmt.Println(Newline)
	fmt.Println("Goal:  140A")
	fmt.Println()
	for y, row := range keypad {
		for x, val := range row {
			pos := Position{x, y}
			if pos.equals(gs.currPosMain) {
				fmt.Print(Underline)
				fmt.Print(Bold)
			}
			if val == "#" {
				fmt.Print(" ")
			} else {
				fmt.Print(val)
			}
			if pos.equals(gs.currPosMain) {
				fmt.Print(Reset)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
	printDir(gs.currPosA)
	fmt.Println()
	printDir(gs.currPosB)
	fmt.Println()
	fmt.Printf("Output: %v\n", gs.output)
	fmt.Println()
}

func printDir(checkPos Position) {
	for y, row := range dirs {
		for x, val := range row {
			pos := Position{x, y}
			if pos.equals(checkPos) {
				fmt.Print(Underline)
				fmt.Print(Bold)
			}
			if val == "#" {
				fmt.Print(" ")
			} else {
				fmt.Print(val)
			}
			if pos.equals(checkPos) {
				fmt.Print(Reset)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
