package main

import "fmt"

func addUser() {
	var input string

	fmt.Printf("What is the name of the user you would like to add? ")
	fmt.Scanln(&input)

	// check if user already exists
	exists := contains(users, input)

	if exists {
		fmt.Println("User already exists.")
	} else {
		users = append(users, userProfile{
			profileName: input,
			groups:      []string{},
			friends:     friendList{nil, 0},
		})
	}
}

func editUser() {
	var input int
	var newName string

	fmt.Println("")

	for i, v := range users {
		fmt.Printf("%v. %v\n", i, v.profileName)
	}

	fmt.Println("Which user would you like to edit?")
	fmt.Scanln(&input)

	fmt.Printf("Change %v to ", users[input].profileName)
	fmt.Scanln(&newName)

	users[input].profileName = newName
}

func deleteUser() {
	var input int

	fmt.Println("")

	for i, v := range users {
		fmt.Printf("%v. %v\n", i+1, v.profileName)
	}

	fmt.Println("Which user would you like to delete?")
	fmt.Scanln(&input)

	copy(users[input-1:], users[input:])
	users = users[:len(users)-1]
}

func contains(s []userProfile, x string) bool {
	for _, v := range s {
		if x == v.profileName {
			return true
		}
	}
	return false
}
