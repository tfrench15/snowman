package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	snowman "github.com/tfrench15/snowman/pkg"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("How many misses do you want to play with?")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading user input")
	}

	misses, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error parsing input to int")
	}

	game := snowman.NewGame()
	game.Play()
}
