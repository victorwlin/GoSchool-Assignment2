package main

import (
	"errors"
	"fmt"
)

func (f *friendList) searchView(friend friend) (input int) {

	for {

		fmt.Println("")

		if friend.lastContact.top == nil {
			fmt.Printf("Friend: %v | Group: %v | Desired Frequency of Contact (in days): %v | Recommended Date of Next Contact: %v\n", friend.name, friend.group, friend.desiredFreqOfContact, "No dates")
		} else {
			fmt.Printf("Friend: %v | Group: %v | Desired Frequency of Contact (in days): %v | Recommended Date of Next Contact: %v\n", friend.name, friend.group, friend.desiredFreqOfContact, friend.lastContact.top.date.AddDate(0, 0, friend.desiredFreqOfContact).Format("02 Jan 2006"))
		}

		fmt.Println("")
		fmt.Println("History of Contact")

		if friend.lastContact.top == nil {
			fmt.Println("No dates")
		} else {
			friend.lastContact.printStack()
		}

		fmt.Println("")
		fmt.Println("1. Add new date of last contact")
		fmt.Println("2. Edit existing date of last contact")
		fmt.Println("3. Edit desired frequency of contact")
		fmt.Println("4. Edit group")
		fmt.Println("5. Edit friend name")
		fmt.Println("6. Delete friend")
		fmt.Println("7. Return to friends menu")
		fmt.Println("8. Exit")
		fmt.Printf("Please select your desired function: ")
		fmt.Scanln(&input)

		// error check
		if input >= 1 && input <= 8 {
			return input
		} else {
			fmt.Println(errors.New("Invalid selection"))
		}

	}
}
