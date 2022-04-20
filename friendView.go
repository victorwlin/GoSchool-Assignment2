package main

import (
	"errors"
	"fmt"
)

func (f *friendList) friendView() (input int) {
	for {
		fmt.Println("")
		fmt.Printf("FRIENDS LIST - User Profile: %v\n", users[user].profileName)

		if f.head == nil {
			fmt.Println("No friends found")
		} else {
			currentFriend := f.head

			// traversal
			for {
				fmt.Println("")

				if currentFriend.lastContact.top == nil {
					fmt.Printf("Friend: %v | Group: %v | Date of Last Contact: %v | Desired Frequency of Contact (in days): %v | Recommended Date of Next Contact: %v\n", currentFriend.name, currentFriend.group, "No dates", currentFriend.desiredFreqOfContact, "No dates")
				} else {
					fmt.Printf("Friend: %v | Group: %v | Date of Last Contact: %v | Desired Frequency of Contact (in days): %v | Recommended Date of Next Contact: %v\n", currentFriend.name, currentFriend.group, currentFriend.lastContact.top.date.Format("02 Jan 2006"), currentFriend.desiredFreqOfContact, currentFriend.lastContact.top.date.AddDate(0, 0, currentFriend.desiredFreqOfContact).Format("02 Jan 2006"))
				}

				if currentFriend.next == nil {
					break
				}

				currentFriend = currentFriend.next
			}
		}

		fmt.Println("")
		fmt.Println("1. Sort list")
		fmt.Println("2. Search for existing friend details and update")
		fmt.Println("3. Add friend")
		fmt.Println("4. Add group")
		fmt.Println("5. Edit existing group")
		fmt.Println("6. Delete group")
		fmt.Println("7. Return to user profiles menu")
		fmt.Println("8. Exit")
		fmt.Printf("Please select your desired function: ")
		fmt.Scanln(&input)

		// error check and return
		if input >= 1 && input <= 8 {
			return input
		} else {
			fmt.Println(errors.New("Invalid selection"))
		}
	}
}
