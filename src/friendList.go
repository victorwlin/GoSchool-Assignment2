package main

type friend struct {
	name                 string
	group                string
	lastContact          *stack
	desiredFreqOfContact int
	next                 *friend
}

type friendList struct {
	head *friend
	size int
}

func (f *friendList) addFriend(name string, group string, lastContact *stack, desiredFreqOfContact int) {
	newFriend := &friend{
		name:                 name,
		group:                group,
		lastContact:          lastContact,
		desiredFreqOfContact: desiredFreqOfContact,
		next:                 nil,
	}

	if f.head == nil {
		f.head = newFriend
	} else {
		currentFriend := f.head

		// traversal
		for currentFriend.next != nil {
			currentFriend = currentFriend.next
		}

		currentFriend.next = newFriend
	}

	f.size++
}

func (f *friendList) removeFriend(i int) {
	if i == 0 {
		f.head = f.head.next
	} else {
		currentFriend := f.head
		prevFriend := f.head
		for index := 0; index < i; index++ {
			prevFriend = currentFriend
			currentFriend = currentFriend.next
		}
		prevFriend.next = currentFriend.next
	}
	f.size--
}

func (f *friendList) traverse(index int) (friend *friend) {
	currentFriend := f.head
	for i := 0; i < index; i++ {
		currentFriend = currentFriend.next
	}
	return currentFriend
}
