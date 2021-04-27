package main

import (
	"fmt"
	"strconv"
)

func main() {

	numbers := []int{230, 711, 13, 37}
	phoneticNumbers(numbers)

}

func phoneticNumbers(nums []int) {
	for i := 0; i < len(nums); i++ {
		s := strconv.Itoa(nums[i]) // Convert decimal into a str
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
		if i < len(nums)-1 { // print commas except last word
			fmt.Printf(",")
		}
	}
}
