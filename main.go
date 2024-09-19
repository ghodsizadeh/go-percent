package main

import (
	"fmt"
	// "time"
)
import (
	"os"
	"strconv"
)

func showProgressBar(percentage int) {
	const barWidth = 50
	progress := int(float64(barWidth) * (float64(percentage) / 100.0))
	fmt.Printf("\r[")
	for i := 0; i < barWidth; i++ {
		if i < progress {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d%%", percentage)
}

func main() {

	// get input as argument
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a number as argument")
	} else {
		// convert string to int
		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid number as argument")
		} else {
			// check if number is between 0 and 100
			if num < 0 || num > 100 {
				fmt.Println("Please provide a number between 0 and 100")
			} else {
				showProgressBar(num)
				fmt.Println()
			}

		}
	}
	// for i := 0; i <= 100; i++ {
	// 	showProgressBar(i)
	// 	time.Sleep(100 * time.Millisecond)
	// }
	// showProgressBar(91)

}
