// Simple Blackjack for CSI 380
// This program play's a simple, single suit game of Blackjack
// against a computer dealer.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Move a card from deck to hand
func drawCard(hand *[]string, deck *[]string) {
	deckSize := len(*deck)
	deck = deck[:deckSize-2]
	hand = append(hand, deck[deckSize-1])
}

// Calculate the score of the hand
func calculateScore(hand []string) int {
	// YOU FILL IN HERE
}

// Print everyone's scores and hands
func printStatus(playerCards, dealerCards []string) {
	// YOU FILL IN HERE
}

// Entry point and main game loop
func main() {
	// YOU FILL IN HERE
	deck := [13]string{ "2" , "3" , "4" , "5" , "6" , "7" , "8" , "9" , "10" , "J", "Q" , "K" ,"A"}
	var playerCards []string
	var dealerCards []string

	//Shuffle randomly the deck
	rand.seed(time.Now().UnixNano())
	//Fisher-Yates Shuffle
	for i := len(deck) - 1 ; i > 0 ; i--
	{
		j := rand.Intn(i+1)
		a[i], a[j] = a[j], a[i]
	}

}
