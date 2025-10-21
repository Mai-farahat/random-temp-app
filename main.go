package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed random generator with current time
	rand.Seed(time.Now().UnixNano())

	// Pick a random temperature between 0 and 40 (inclusive)
	temperature := rand.Intn(41) // 0..40

	// Print temperature
	fmt.Printf("Temperature: %d°C\n", temperature)

	// Switch statement to print message
	switch {
	case temperature < 10:
		fmt.Println("It's cold 🥶.")
	case temperature <= 25:
		fmt.Println("It's moderate 🍀.")
	default:
		fmt.Println("It's warm 😎.")
	}
}