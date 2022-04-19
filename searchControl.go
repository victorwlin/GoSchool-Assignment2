package main

func searchControl() {
	sCon := true

	friend, i := users[user].friends.seqSearch()

	for sCon {
		input := users[user].friends.searchView(*friend)

		switch input {
		case 1:
			addDate(*friend)
		case 2:
			editDate(*friend)
		case 3:
			editDesiredFreq(friend)
		case 4:
			editGroup(friend)
		case 5:
			editName(friend)
		case 6:
			users[user].friends.removeFriend(i)
			sCon = false
		case 7:
			sCon = false
		case 8:
			sCon = false
			fCon = false
			exitCon = false
		}
	}
}
