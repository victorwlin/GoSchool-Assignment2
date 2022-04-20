package main

import (
	"errors"
	"fmt"
)

func (f *friendList) sortControl() {
	input := sortView()

	for i := 1; i < f.size; i++ {
		// copy
		copy := f.traverse(i)

		// loop over sorted and shift to right as necessary
		sorted := i
		switch input {

		// friend name
		case 1:
			for sorted > 0 && copy.name < f.traverse(sorted-1).name {
				f.traverse(sorted - 1).next = f.traverse(sorted + 1)
				copy.next = f.traverse(sorted - 1)

				if sorted <= 1 {
					f.head = copy
				} else {
					f.traverse(sorted - 2).next = copy
				}

				sorted--
			}

		// group
		case 2:
			for sorted > 0 && copy.group < f.traverse(sorted-1).group {
				f.traverse(sorted - 1).next = f.traverse(sorted + 1)
				copy.next = f.traverse(sorted - 1)

				if sorted <= 1 {
					f.head = copy
				} else {
					f.traverse(sorted - 2).next = copy
				}

				sorted--
			}

		// date of last contact
		case 3:
			for sorted > 0 && copy.lastContact.top.date.After(f.traverse(sorted-1).lastContact.top.date) {
				f.traverse(sorted - 1).next = f.traverse(sorted + 1)
				copy.next = f.traverse(sorted - 1)

				if sorted <= 1 {
					f.head = copy
				} else {
					f.traverse(sorted - 2).next = copy
				}

				sorted--
			}

		// desired frequency of contact
		case 4:
			for sorted > 0 && copy.desiredFreqOfContact < f.traverse(sorted-1).desiredFreqOfContact {
				f.traverse(sorted - 1).next = f.traverse(sorted + 1)
				copy.next = f.traverse(sorted - 1)

				if sorted <= 1 {
					f.head = copy
				} else {
					f.traverse(sorted - 2).next = copy
				}

				sorted--
			}

		// recommended date of next contact
		case 5:
			for sorted > 0 && copy.lastContact.top.date.AddDate(0, 0, copy.desiredFreqOfContact).Before(f.traverse(sorted-1).lastContact.top.date.AddDate(0, 0, f.traverse(sorted-1).desiredFreqOfContact)) {
				f.traverse(sorted - 1).next = f.traverse(sorted + 1)
				copy.next = f.traverse(sorted - 1)

				if sorted <= 1 {
					f.head = copy
				} else {
					f.traverse(sorted - 2).next = copy
				}

				sorted--
			}

		default:
			fmt.Println(errors.New("Error in sortControl()"))
		}
	}

}
