package main

import (
	"errors"
	"fmt"
)

func (f *friendList) friendControl() {
	for fCon {
		input := users[user].friends.friendView()

		switch input {
		case 1:
			f.sortControl()

		case 2:
			searchControl()

		case 3:
			users[user].friends.addFriendToList()

		case 4:
			addGroup()

		case 5:
			editExistingGroup()

		case 6:
			deleteGroup()

		case 7:
			fCon = false

		case 8:
			fCon = false
			uCon = false
			exitCon = false

		default:
			fmt.Println(errors.New("Error in friendControl()"))
		}
	}
}
