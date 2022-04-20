package main

import (
	"errors"
	"fmt"
	"strconv"
)

func userView() (input string) {
	for {
		fmt.Println("")
		fmt.Println("Welcome to Friend Tracker, an app that keeps you in touch with your friends!")
		fmt.Println("")
		fmt.Println("USER PROFILES AND FUNCTIONS")

		for i, v := range users {
			fmt.Printf("%v. %v\n", i+1, v.profileName)
		}

		fmt.Println("a. Add new user profile")
		fmt.Println("b. Edit existing user profile")
		fmt.Println("c. Delete user profile")
		fmt.Println("z. Exit")
		fmt.Printf("Please select your user profile or desired function: ")
		fmt.Scanln(&input)

		// error check and return
		if (input >= "1" && input <= strconv.Itoa(len(users))) || input == "a" || input == "b" || input == "c" || input == "z" {
			return input
		} else {
			fmt.Println(errors.New("Invalid selection"))
		}
	}
}
