package snowman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

// Game represents the game of Snowman, including the secret word,
// number of misses the player is allowed, and fields for tracking
// the guesses so far.
type Game struct {
	secret         string
	missesAllowed  int
	currentGuesses map[string]bool
	board          []string
}

// NewGame is a constructor returning a new instance of Game.
func NewGame(words []string, misses int) *Game {
	m := make(map[string]bool)
	b := []string{}

	idx := rand.Intn(len(words))

	return &Game{
		secret:         words[idx],
		missesAllowed:  misses,
		currentGuesses: m,
		board:          b,
	}
}

// Play sets up and executes the main game loop and continues looping
// until the game is won or lost.
func (g *Game) Play() {
	for {
		g.printBoard()

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please guess again")
			continue
		}

		if !unicode.IsLetter(rune(input)) {
			fmt.Println("Guess must be a letter, please guess again")
			continue
		}

		guess := strings.ToLower(strings.TrimSpace(input))
		if len(guess) != 1 {
			fmt.Println("Guess must be one letter, please guess again")
			continue
		}

		g.processGuess(guess)
		if g.isWon() {
			fmt.Printf("Congratulations, the secret word was %s you won!", g.secret)
			break
		}
		if g.isLost() {
			fmt.Printf("Oh no, you lost! The secret word was %s", g.secret)
			break
		}
	}
}

func (g *Game) processGuess(guess string) {
	if strings.Contains(g.secret, guess) {
		for i, ch := range g.secret {
			if string(ch) == guess {
				g.board[i] = guess
			}
		}
		return
	}

	g.currentGuesses[guess] = true
}

func (g *Game) isWon() bool {
	board := strings.Join(g.board, "")
	if strings.ContainsAny(board, "_") {
		return false
	}

	return true
}

func (g *Game) isLost() bool {
	if len(g.currentGuesses) == g.missesAllowed {
		return true
	}

	return false
}

func (g *Game) printBoard() {
	fmt.Printf("Current board: %v", g.board)

	var lst []string
	for key := range g.currentGuesses {
		lst = append(lst, key)
	}
	fmt.Printf("Guesses so far: %v", lst)
	fmt.Printf("You have %d guesses remaining", g.missesAllowed-len(g.currentGuesses))
}
