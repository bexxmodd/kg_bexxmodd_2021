package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // discard program name
	var err error
	numbers := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		if numbers[i], err = strconv.Atoi(args[i]); err != nil {
			fmt.Println("Error: Program only accepts integers")
			panic(err)
		}
	}
	printPhonetic(numbers)
}

// printPhonetic: Takes array of integers and prints phonetic equivalent
//  of each integer into the stdout seperated with commas.
// Example:
//		- {25, 0, 313} will be printed as -> "TwoFive,Zero,ThreeOneThree"
func printPhonetic(nums []int) {
	for i := 0; i < len(nums); i++ {
		s := strconv.Itoa(nums[i]) // convert decimal into a str
		for j := 0; j < len(s); j++ {

			switch digit := string(s[j]); digit {
			case "0":
				fmt.Printf("Zero")
			case "1":
				fmt.Printf("One")
			case "2":
				fmt.Printf("Two")
			case "3":
				fmt.Printf("Three")
			case "4":
				fmt.Printf("Four")
			case "5":
				fmt.Printf("Five")
			case "6":
				fmt.Printf("Six")
			case "7":
				fmt.Printf("Seven")
			case "8":
				fmt.Printf("Eight")
			default:
				fmt.Printf("Nine")
			}
		}
		// print commas; except after the last word
		if i < len(nums)-1 {
			fmt.Printf(",")
		}
	}
}
