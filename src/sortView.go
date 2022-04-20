package main

import (
	"errors"
	"fmt"
)

func sortView() (input int) {
	for {
		fmt.Println("")
		fmt.Println("By which field would you like to sort the list?")
		fmt.Println("1. Friend name")
		fmt.Println("2. Group")
		fmt.Println("3. Date of last contact")
		fmt.Println("4. Desired frequency of contact")
		fmt.Println("5. Recommended date of next contact")
		fmt.Printf("Selection: ")
		fmt.Scanln(&input)

		// error check
		if input >= 1 && input <= 5 {
			return input
		} else {
			fmt.Println(errors.New("Invalid selection"))
		}
	}
}
