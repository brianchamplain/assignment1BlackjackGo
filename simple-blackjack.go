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

//used in calculate
const valueMap map[string]int = {
	"2" : 2,
	"3" : 3,
	"4" : 4,
	"5" : 5,
	"6" : 6,
	"7" : 7,
	"8" : 8,
	"9" : 9,
	"10":10,
	"J" :10,
	"Q" :10,
	"K" :10,
	"A" :11 //is checked in calculate score directly
}

// Move a card from deck to hand
func drawCard(hand *[]string, deck *[]string) {
	hand = append(hand, deck[deckSize-1]) //add last element from deck slice to the end of the hand
	deck = deck[:len(*deck)-2] //should remove the last element from the deck
}

// Calculate the score of the hand
func calculateScore(hand []string) int {
	sum := 0
	aces := 0 //track aces and subtract them from the sum after if above 21
	for i := 0; i < len(hand); i++ {
		sum += valueMap[hand[i]] //get point value from valuemap constant defined above
		if(hand[i] == "A") {
			aces++ //count them and subtract later if needed
		}
	}
	for sum > 21 && aces > 0 { //if the score is above 21 and an ace was in the hand
		sum -= 10
		aces--
	}
}

// Print everyone's scores and hands
func printStatus(playerCards, dealerCards []string) {
	fmt.Println("Player Cards:")
	for i:= 0; i < len(playerCards); i++ { //print all player cards
		fmt.Println(playerCards[i])
	}
	fmt.Println("Player Score: " + strconv.Itoa(calculateScore(playerCards)))
	fmt.Println("Dealer Cards:")
	for i:= 0; i < len(dealerCards); i++ { //print all dealer cards
		fmt.Println(dealerCards[i])
	}
	fmt.Println("Dealer Score: " + strconv.Itoa(calculateScore(dealerCards)))
}

// Entry point and main game loop
func main() {
	// YOU FILL IN HERE
deck: = [13]string{ "2" , "3" , "4" , "5" , "6" , "7" , "8" , "9" , "10" , "J", "Q" , "K" ,"A" }
	var playerCards[]string
	var dealerCards[]string

	//Shuffle randomly the deck
	rand.seed(time.Now().UnixNano())
	//Fisher-Yates Shuffle
	for i : = len(deck) - 1; i > 0; i--
	{
	j: = rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	  fmt.Printf("Dealer draws first card.")
		  drawCard(dealerCards, deck)
		  fmt.Printf("Player receives two cards.")
		  drawCard(playerCards, deck)
		  drawCard(playerCards, deck)
		  printStatus(playerCards, dealerCards)

		  //player decision loop
		  var decision bool = true
		  for decision == true {
			  fmt.Printf("Do you want to (H)it, (S)tay, or (Q)uit?")
				  var response int
				  fmt.Scanf("%c", &response)
				  if response == "H" || response == "h" {
					  drawCard(playerCards, deck)
						  printStatus(playerCards, dealerCards)
						  if calculateScore(playerCards) > 21{
							  fmt.Printf("You busted! You lose!")
								  os.exit(1)
						  }
				  }
				  else if response == "S" || response == "s"{
					  fmt.Printf("You chose to stay")
						  printStatus(playerCards, dealerCards)
				  }
				  else {
					  os.exit(1)
				  }

		  }

	  //Dealer's turn
	  fmt.Printf("Dealer draws rest of cards.")
		  for calculateScore(dealerCards) < 17	{
			  drawCard(dealerCards, deck)
		  }
	  printStatus(playerCards, dealerCards)

		  if calculateScore(dealerCards) > 21 {
			  fmt.Printf("Dealer busts! You win!")
		  }
		  else if calculateScore(dealerCards) < calculateScore(playerCards) {
			  fmt.Printf("Dealer wins!")
		  }
		  else {
			  fmt.Printf("It's a tie!")
		  }
	os.exit(1)
}
