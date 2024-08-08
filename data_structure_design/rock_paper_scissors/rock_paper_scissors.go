package main

import (
	"fmt"
	"math/rand"
)

const (
	rock         = "rock"
	paper        = "paper"
	scissors     = "scissors"
	playerWins   = "Player wins"
	computerWins = "Computer wins"
	tied         = "Tied game"
)

func main() {
	choice := ""
	resultMatrix := map[string]map[string]string{
		"rock": {
			"rock":     tied,
			"paper":    computerWins,
			"scissors": playerWins,
		},
		"paper": {
			"rock":     playerWins,
			"paper":    tied,
			"scissors": computerWins,
		},
		"scissors": {
			"rock":     computerWins,
			"paper":    playerWins,
			"scissors": tied,
		},
	}
	computerOptions := []string{rock, paper, scissors}

	for choice != "exit" {
		fmt.Println("Choose your move (rock, paper, scissors, exit): ")
		fmt.Scanln(&choice)

		computerChoice := getComputerChoice(computerOptions)
		results := resultMatrix[choice][computerChoice]
		fmt.Printf("Player: %s, Computer: %s, Result: %s \n", choice, computerChoice, results)
	}
}

func getComputerChoice(computerOptions []string) string {
	return computerOptions[rand.Intn(len(computerOptions))]
}
