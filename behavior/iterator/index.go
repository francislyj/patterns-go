package iterator

import "fmt"

func TestIterator() {
	user1 := &user{
		name: "aa",
		age: 11,
	}

	user2 := &user{
		name: "bb",
		age: 12,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("user name %s age %d \n", user.name, user.age)
	}
}
