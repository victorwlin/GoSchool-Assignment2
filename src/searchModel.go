package main

import (
	"errors"
	"fmt"
	"time"
)

func (f *friendList) seqSearch() (friend *friend, i int) {
	var name string

	fmt.Println("")
	fmt.Printf("Search friends by name: ")
	fmt.Scanln(&name)

	i = 0
	currentFriend := f.head
	for {
		if currentFriend.name == name {
			return currentFriend, i
		} else if currentFriend.next == nil {
			fmt.Println(errors.New("Friend not found."))
			return currentFriend.next, -1
		} else {
			currentFriend = currentFriend.next
			i++
		}
	}
}

func addDate(friend friend) {
	var (
		year  uint
		month uint
		date  uint
	)

	fmt.Println("")
	fmt.Printf("What year did you last make contact? ")
	fmt.Scanln(&year)
	fmt.Printf("What month did you last make contact? (Enter month as an integer) ")
	fmt.Scanln(&month)
	fmt.Printf("What date did you last make contact? ")
	fmt.Scanln(&date)

	friend.lastContact.push(time.Date(int(year), time.Month(month), int(date), 11, 0, 0, 0, time.UTC))
}

func editDate(friend friend) {
	var (
		year  uint
		month uint
		date  uint
	)

	fmt.Println("")
	fmt.Printf("What year did you last make contact? ")
	fmt.Scanln(&year)
	fmt.Printf("What month did you last make contact? (Enter month as an integer) ")
	fmt.Scanln(&month)
	fmt.Printf("What date did you last make contact? ")
	fmt.Scanln(&date)

	friend.lastContact.top.date = time.Date(int(year), time.Month(month), int(date), 11, 0, 0, 0, time.UTC)
}

func editDesiredFreq(friend *friend) {
	var desiredFreq int

	fmt.Println("")
	fmt.Printf("How often would you like to make contact? (in days) ")
	fmt.Scanln(&desiredFreq)

	friend.desiredFreqOfContact = desiredFreq
}

func editGroup(friend *friend) {
	var group int

	fmt.Println("")
	fmt.Println("Please select what group you would like this friend to belong to.")

	for i, v := range users[user].groups {
		fmt.Printf("%v. %v\n", i, v)
	}

	fmt.Printf("Selection: ")
	fmt.Scanln(&group)

	friend.group = users[user].groups[group]
}

func editName(friend *friend) {
	var name string

	fmt.Printf("What would you like to change this friend's name to? ")
	fmt.Scanln(&name)

	friend.name = name
}
