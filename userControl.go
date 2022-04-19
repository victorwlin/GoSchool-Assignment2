package main

import (
	"errors"
	"fmt"
	"strconv"
)

func userControl() {
	for uCon {
		input := userView()

		// if user selected a user profile
		if input >= "1" && input <= strconv.Itoa(len(users)) {
			user, _ = strconv.Atoi(input)
			user = user - 1
			fCon = true
			users[user].friends.friendControl()
		} else {
			switch input {
			case "a":
				addUser()
			case "b":
				editUser()
			case "c":
				deleteUser()
			case "z":
				uCon = false
				exitCon = false
			default:
				fmt.Println(errors.New("Error in userControl()"))
			}
		}
	}
}
