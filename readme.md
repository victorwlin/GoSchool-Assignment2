#### Welcome to Friend Tracker, an app that keeps you in touch with your friends!
This assignment was submitted for the GoSchool Go Advanced module.

#### Features
- The application begins on the User Menu, where you can select your user profile. Please feel free to create a new user profile, edit an existing user profile, or delete a user profile.
- After you have selected a user profile, you will be taken to the Friend Menu, where you can see a list of all of the friends that you have added to the application. Information displayed includes name, group, date of last contact, desired frequency of contact (in days), and recommended date of next contact.
- From the Friend Menu, you can sort your friends, search your list for a specific friend, add a friend, add a group, edit an existing group, or delete a group.
- Please feel free to sort your friends according to name, group, date of last contact, desired frequency of contact, or recommended date of next contact.
- To get a more detailed history of contact or to modify any details about your friend, please do a search. This will take you to the Search Menu.
- From the Search Menu, you can add a new date of last contact, edit the existing date of last contact, edit desired frequency of contact, edit friend’s group, edit friend name, or delete the friend.

#### Data Structures
##### Linked List
- Each user profile has an associated Linked List with all friend data.
- Linked List was chosen for friends because:
  - The amount of friends that people have varies from person to person. The size of a Linked List is not fixed and can grow as large as is necessary.
  - I can add, delete, and sort by updating pointers rather than shifting around data.
  - Other than some space for pointers, no storage space has been wasted.
- The search friends function uses sequential search.
- The sort friends function uses insertion sort by manipulating pointers.

##### Stack
- Each friend in the Linked List has a Stack data structure with an entire history of contact.
- Stack was chosen for history of contact because:
  - The most important date, the date of last contact, is popped onto the top of the Stack and is easy/fast to access. This is the date that is used in the calculation of the recommended date of next contact.
  - As the user makes contact with his/her friends, he/she can simply pop the date of last contact onto the Stack.
  - Users are allowed to edit the date of last contact on the top of the Stack, but I did not intend for them to be allowed to edit the entire history of contact. Once the history is created, it is meant to be permanent.

#### Data Format
- Data is centered around a userProfile struct with a profileName string, a slice of strings denoting groups, and a Linked List containing all friends.
- Each friend node of the Linked List contains a name string, group string, a Stack with all dates of contact, desiredFreqOfContact int, and a pointer to the next friend node.

#### Application Organization
- The application was built with the Model-View-Controller (MVC) design pattern.
- Each major menu has a View file that displays the menu to the user and receives input. The Controller file receives input from the View file and executes the appropriate function in the Model file.
- Major menus that have MVC files include User Menu, Friend Menu, Sort Menu, and Search Menu.
- Specialized data structures (Linked List and Stack) have their own file with a few methods defined for manipulation.
- All data is initialized in the data.go file.

#### Error Handling
- Error handling is done mostly in the View files of each section. The user input is checked to make sure it makes sense, and if it is invalid, the user will be asked to re-enter his/her choice.
- Each Control file also has a default case in the switch statement, just in case what was passed to Control from View is invalid.
- The main.go file also has a panic recovery deferred function to print any unexpected errors that occur.

#### Concurrency
- Concurrency is used in the friendModel.go file, specifically the editExistingGroup and deleteGroup functions.
- editExistingGroup uses goroutines to concurrently run two functions: one to update the group name of all friends who currently belong to the group and the other to update the group name in the group slice.
- deleteGroup makes use of three goroutines.
  - The first goroutine identifies all friends belonging to this group and sends a slice of the index position of each friend that needs to be deleted to the deleteList channel.
  - The second goroutine receives the slice from the deleteList channel and deletes each of the friends.
  - The third goroutine deletes the group from the groups slice.
