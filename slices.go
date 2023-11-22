package main

import "fmt"

// arrays
func getMessagesWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func arrayTest() {
	messages := getMessagesWithRetries()
	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i])
	}
}

// slices
func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		costs[i] = float64(len(message)) * 0.01
	}
	return costs
}

// variadic functions
// it can take any number of integers
func sum(nums ...int) int {
	// fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	// fmt.Println(total)
	return total
}

func createMatrix(rows, cols int) [][]int {
	// create a slice of slices
	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func sliceTest() {
	// messages := []string{
	// 	"click here to sign up",
	// 	"pretty please click here",
	// 	"we beg you to sign up",
	// }
	// costs := getMessageCosts(messages)
	// fmt.Println(costs)

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	// spread operator
	// pass in a slice of ints
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))

	// create a 2d matrix
	matrix := createMatrix(10, 10)
	fmt.Println(matrix)
}
