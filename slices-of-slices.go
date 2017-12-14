package main

import (
	"fmt"
	"strings"
)

func main() {
	game := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
