package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 100
	secretNum := rand.Intn(n)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guess the number!")
	fmt.Printf("Please enter a number between 1 and %d:", n)
	var input int
	for {
		fmt.Scanf("%d", &input)
		if input == secretNum {
			fmt.Println("You guessed it!")
			break
		} else if input < secretNum {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}
}
