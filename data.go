package main

import "time"

var (
	user    int
	exitCon bool = true
	uCon    bool = true
	fCon    bool = true
)

type userProfile struct {
	profileName string
	groups      []string
	friends     friendList
}

var users = []userProfile{
	{
		profileName: "Victor",
		groups:      []string{"USA", "SG"},
		friends:     friendList{nil, 0},
	},
	{
		profileName: "Tokey",
		groups:      []string{"College", "Work"},
		friends:     friendList{nil, 0},
	},
}

func init() {
	jimmyStack := stack{nil, 0}
	jimmyStack.push(time.Date(2022, 1, 11, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))
	jimmyStack.push(time.Date(2022, 4, 14, 11, 0, 0, 0, time.UTC))

	benStack := stack{nil, 0}
	benStack.push(time.Date(2022, 1, 31, 11, 0, 0, 0, time.UTC))
	benStack.push(time.Date(2022, 2, 6, 11, 0, 0, 0, time.UTC))

	angelStack := stack{nil, 0}
	angelStack.push(time.Date(2022, 4, 9, 11, 0, 0, 0, time.UTC))
	angelStack.push(time.Date(2022, 4, 13, 11, 0, 0, 0, time.UTC))

	arjunStack := stack{nil, 0}
	arjunStack.push(time.Date(2022, 3, 29, 11, 0, 0, 0, time.UTC))
	arjunStack.push(time.Date(2022, 4, 14, 11, 0, 0, 0, time.UTC))

	users[user].friends.addFriend("Jimmy", "USA", &jimmyStack, 30)
	users[user].friends.addFriend("Ben", "USA", &benStack, 60)
	users[user].friends.addFriend("Angel", "SG", &angelStack, 60)
	users[user].friends.addFriend("Arjun", "SG", &arjunStack, 30)
}
