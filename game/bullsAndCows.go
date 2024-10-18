package game

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

const (
	GameRunning = 1
	GameEnded   = 2
)

type Guess struct {
	Word   string
	MadeAt time.Time
}

type Game struct {
	Id        string
	startTime time.Time
	word      string
	wordLen   int
	state     int
	Guesses   []*Guess
}

type MatchResponse struct {
	Valid   bool
	Exact   bool
	Bulls   int
	Cows    int
	Message string
}

func NewGame() *Game {
	return &Game{
		Id:        uuid.NewString(),
		startTime: time.Now().UTC(),
		word:      getRandomWord(),
		wordLen:   4,
		state:     GameRunning,
		Guesses:   make([]*Guess, 0),
	}
}

func (g *Game) Validate(guess string) error {
	if g.state != GameRunning {
		return errors.New("game not running")
	}
	if g.wordLen != len(guess) {
		return errors.New("invalid guess length")
	}
	return nil
}

func (g *Game) Match(guess string) *MatchResponse {
	b := 0
	c := 0
	if err := g.Validate(guess); err != nil {
		return &MatchResponse{
			Valid:   false,
			Exact:   false,
			Bulls:   0,
			Cows:    0,
			Message: err.Error(),
		}
	}
	g.Guesses = append(g.Guesses, &Guess{
		Word:   guess,
		MadeAt: time.Now().UTC(),
	})

	wordCount := make(map[rune]int)
	guessCount := make(map[rune]int)

	for pos, char := range guess {
		if rune(g.word[pos]) == char {
			b++
		} else {
			wordCount[rune(g.word[pos])]++
			guessCount[char]++
		}
	}

	for char, count := range guessCount {
		if wordCount[char] > 0 {
			c += min(count, wordCount[char])
		}
	}

	exact := b == g.wordLen
	message := "Carry on"
	if exact {
		g.state = GameEnded
		message = "Victory"

	}
	return &MatchResponse{
		Valid:   true,
		Exact:   exact,
		Bulls:   b,
		Cows:    c,
		Message: message,
	}
}

func getRandomWord() string {
	words := []string{
		"able", "acid", "aged", "also", "area", "army", "away", "baby", "back", "ball",
		"band", "bank", "base", "bath", "bear", "beat", "been", "beer", "bell", "belt",
		"best", "bill", "bird", "blow", "blue", "boat", "body", "bomb", "bond", "bone",
		"book", "boom", "born", "boss", "both", "bowl", "bulk", "burn", "bush", "busy",
		"call", "calm", "camp", "card", "care", "case", "cash", "cast", "cell", "chat",
		"chip", "city", "club", "coal", "coat", "code", "cold", "come", "cook", "cool",
		"cope", "copy", "core", "cost", "crew", "crop", "dark", "data", "date", "dawn",
		"days", "dead", "deal", "dear", "debt", "deep", "deny", "desk", "dial", "diet",
		"disc", "does", "done", "door", "dose", "down", "draw", "drop", "drug", "dual",
		"duke", "dust", "duty", "each", "earn", "ease", "east", "easy", "edge", "else",
	}
	return words[rand.Intn(len(words))]
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
