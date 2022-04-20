package main

import (
	"fmt"
	"runtime"
	"sync"
)

func (f *friendList) addFriendToList() {
	var (
		name                 string
		group                int
		desiredFreqOfContact int
	)

	fmt.Println("")
	fmt.Println("Please select group to add friend to.")

	for i, v := range users[user].groups {
		fmt.Printf("%v. %v\n", i+1, v)
	}

	fmt.Printf("Selection: ")
	fmt.Scanln(&group)

	fmt.Printf("Friend name: ")
	fmt.Scanln(&name)

	fmt.Printf("Desired frequency of contact (in days): ")
	fmt.Scanln(&desiredFreqOfContact)

	users[user].friends.addFriend(name, users[user].groups[group-1], &stack{nil, 0}, desiredFreqOfContact)
}

func addGroup() {
	for {
		var group string

		fmt.Println("")
		fmt.Printf("What group would you like to add? ")
		fmt.Scanln(&group)

		// error checking
		exists := false
		for _, v := range users[user].groups {
			if v == group {
				exists = true
			}
		}

		if exists {
			fmt.Println("Group already exists.")
		} else {
			users[user].groups = append(users[user].groups, group)
			fmt.Println("Group has been added.")
			break
		}
	}
}

func editExistingGroup() {
	var group int
	var newName string
	var oldName string

	fmt.Println("")
	for i, v := range users[user].groups {
		fmt.Printf("%v. %v\n", i+1, v)
	}
	fmt.Printf("Please select group you would like to edit: ")
	fmt.Scanln(&group)
	oldName = users[user].groups[group-1]

	fmt.Printf("What would you like to change the group name to? ")
	fmt.Scanln(&newName)

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	// update all friends that are part of the changed group
	go func() {
		defer wg.Done()

		currentFriend := users[user].friends.head
		for {
			if currentFriend.group == oldName {
				currentFriend.group = newName
			}

			if currentFriend.next == nil {
				break
			} else {
				currentFriend = currentFriend.next
			}
		}
	}()

	// change group in the group slice
	go func() {
		defer wg.Done()

		users[user].groups[group-1] = newName
	}()

	wg.Wait()
}

func deleteGroup() {
	var i int

	fmt.Println("")
	for i, v := range users[user].groups {
		fmt.Printf("%v. %v\n", i+1, v)
	}
	fmt.Printf("Please select group you would like to delete: ")
	fmt.Scanln(&i)
	group := users[user].groups[i-1]

	// this channel was added for practice
	// it is definitely not required nor helpful in this situation
	deleteList := make(chan []int)

	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	wg.Add(3)

	// identify all occurences
	go func() {
		defer wg.Done()

		tempList := []int{}

		currentFriend := users[user].friends.head
		for index := 0; index < users[user].friends.size; index++ {
			if currentFriend.group == group {
				tempList = append(tempList, index)
			}
			currentFriend = currentFriend.next
		}

		deleteList <- tempList
		close(deleteList)
	}()

	// delete all identified friends of the group
	go func() {
		defer wg.Done()

		tempList, ok := <-deleteList

		if !ok {
			return
		} else {
			for i, v := range tempList {
				users[user].friends.removeFriend(v - i)
			}
		}

	}()

	// // delete all friends in group
	// go func() {
	// 	defer wg.Done()

	// 	// identify all occurences
	// 	deleteList := []int{}
	// 	currentFriend := users[user].friends.head
	// 	for index := 0; index < users[user].friends.size; index++ {
	// 		if currentFriend.group == group {
	// 			deleteList = append(deleteList, index)
	// 		}
	// 		currentFriend = currentFriend.next
	// 	}

	// 	// delete friends
	// 	for i, v := range deleteList {
	// 		users[user].friends.removeFriend(v - i)
	// 	}
	// }()

	// delete group from group slice
	go func() {
		defer wg.Done()

		copy(users[user].groups[i-1:], users[user].groups[i:])
		users[user].groups = users[user].groups[:len(users[user].groups)-1]
	}()

	wg.Wait()
}
