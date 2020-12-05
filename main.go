package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/namsral/flag"
)

const (
	defaultNumCards  = 8
	defaultNumGames  = 1
	defaultVerbosity = false
)

var (
	numCards  int
	numGames  int
	verbosity bool
)

func init() {
	rand.Seed(time.Now().Unix())
}

type gameLog struct {
	verbosity bool
}

func (g gameLog) log(msg string) {
	if g.verbosity {
		fmt.Println(msg)
	}
}

func removeFromCards(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func getDecision(deckSelection, player1Selection, player2Selection int) (int, int) {
	player1Score := 0
	player2Score := 0

	if player1Selection > player2Selection {
		player1Score = deckSelection
	} else if player1Selection < player2Selection {
		player2Score = deckSelection
	} else {
		player1Score = deckSelection
		player2Score = deckSelection
	}
	return player1Score, player2Score
}

type Pack struct {
	score int
	cards []int
}

func createPack(initScore int, initCards []int) Pack {
	return Pack{score: initScore, cards: initCards}
}

func (p *Pack) RandomSelectFromPack() int {
	randIndex := rand.Intn(len(p.cards))
	selection := p.cards[randIndex]
	p.cards = removeFromCards(p.cards, randIndex)
	return selection
}

func (p *Pack) EqualSelectFromPack(cardToSelect int) int {
	found := -1
	for index, card := range p.cards {
		if card == cardToSelect {
			found = index
		}
	}

	if found != -1 {
		selection := p.cards[found]
		p.cards = removeFromCards(p.cards, found)
		return selection
	}

	panic("Why are we still here just to suffer?")
}

func (p *Pack) AddScore(value int) {
	p.score += value
}

func hitMe(totalCards int) (Pack, Pack, Pack) {
	finalPacks := []Pack{}
	for i := 0; i < 3; i++ {
		cards := []int{}
		for j := 0; j < totalCards; j++ {
			cards = append(cards, j+1)
		}

		finalPacks = append(finalPacks, createPack(0, cards))
	}

	return finalPacks[0], finalPacks[1], finalPacks[2]
}

func game(totalCards int, gameLogger gameLog) int {
	totalDeckCards := totalCards
	deckCards, player1, player2 := hitMe(totalDeckCards)

	for turnCount := 0; turnCount < totalDeckCards; turnCount++ {
		gameLogger.log(fmt.Sprintf("Round %d", turnCount))

		deckSelection := deckCards.RandomSelectFromPack()
		gameLogger.log(fmt.Sprintf("Deck Card is %d", deckSelection))

		player1Selection := player1.RandomSelectFromPack()
		player2Selection := player2.EqualSelectFromPack(deckSelection)

		gameLogger.log(fmt.Sprintf("Player1 Chose: %d; Player2 Chose: %d", player1Selection, player2Selection))

		player1Score, player2Score := getDecision(deckSelection, player1Selection, player2Selection)

		player1.AddScore(player1Score)
		player2.AddScore(player2Score)

		gameLogger.log(fmt.Sprintf("Player1 score: %d; Player2 score: %d", player1.score, player2.score))
	}

	gameLogger.log(fmt.Sprintf("Final score : Player1 == %d; Player2 == %d", player1.score, player2.score))

	winner := 0
	// If draw then don't add to count
	if player1.score > player2.score {
		winner = 1
	} else if player1.score < player2.score {
		winner = 2
	}

	return winner
}

func main() {
	flag.IntVar(&numCards, "cards", defaultNumCards, "The number of cards to start in each deck")
	flag.IntVar(&numGames, "games", defaultNumGames, "The number of games to run")
	flag.BoolVar(&verbosity, "verbose", defaultVerbosity, "Toggle if game logs should be shown")
	flag.Parse()

	gameLogger := gameLog{verbosity: verbosity}

	player1Count := 0
	player2Count := 0

	for i := 0; i < numGames; i++ {
		winner := game(numCards, gameLogger)
		// If draw then don't add to count
		if winner == 1 {
			player1Count += 1
		} else if winner == 2 {
			player2Count += 1
		}

		gameLogger.log(fmt.Sprintf("Winner is Player %d", winner))
	}

	fmt.Printf("Player1 Count : %d; Player2 Count: %d; P1's win probability %f%%\n", player1Count, player2Count, (float64(player1Count)/float64(numGames))*100)
}
