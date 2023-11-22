package main

import "fmt"

// regular for loop
func bulkSend(numbers int) float64 {
	totalCost := 0.0
	for i := 0; i < numbers; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

// loops in go can omit some conditions
// for example, this is an infinite loop
func maxMessages(thread float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thread {
			return i
		}
	}
}

// while loop in go
// just a for loop that only has a condition
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	acutalCostInPennies := 1.0
	maxMessagesToSend := 0
	for acutalCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		acutalCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

func loopTest() {
	fmt.Println(bulkSend(10))
	fmt.Println(maxMessages(10.0))
	fmt.Println(maxMessages(20.0))
	fmt.Println(maxMessages(40.0))
	fmt.Println(getMaxMessagesToSend(1.01, 100))
}
