package main

import (
	"fmt"
	"math/rand"
)

var turnScore int = 0
var totalScore int = 0
var turns int = 0
var choice string

func hold(tscore int) {
	totalScore += tscore
	fmt.Println("Your turn score is ", turnScore, "and your total score is", totalScore)
	turnScore = 0
	turn()
}

func roll() {
	turns++
	fmt.Println("---------TURN : ", turns, "---------")
	rolledNum := rand.Intn(5) + 1
	fmt.Println("You Rolled : ", rolledNum)
	if rolledNum == 1 {
		fmt.Println("Turn Over. No Score")
		turnScore = 0
		hold(turnScore)
	} else {
		turnScore += rolledNum
		if (turnScore + totalScore) >= 20 {
			fmt.Println("\n!!!!!-----You Win-----!!!!!\nYou finished in", turns, "turns with", turnScore+totalScore, "points")
		} else {
			fmt.Println("Your turn score is ", turnScore, "and your total score is", totalScore)
			fmt.Println("If you hold, you will have", turnScore+totalScore, "points")
			fmt.Println("Enter 'r' to roll again or 'h' to hold")
			fmt.Scan(&choice)
			switch choice {
			case "r":
				roll()
			case "h":
				hold(turnScore)
			}
		}
	}
}

func turn() {
	fmt.Println("\nWelcome to the game of Pig!")
	fmt.Println("Enter 'r' to roll")
	fmt.Scan(&choice)
	roll()
}

func main() {
	fmt.Println("\nLet's Play PIG!")
	fmt.Println("\n---------RULES---------")
	fmt.Println("* See how many turns it takes you to get to 20.")
	fmt.Println("* Turn ends when you hold or roll a 1.")
	fmt.Println("* If you roll a 1, you lose all points for the turn.")
	fmt.Println("* If you hold, you save all points for the turn.")
	fmt.Print("\nLet's Start\n")
	turn()
}
