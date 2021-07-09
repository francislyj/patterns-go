package snapshot

import "fmt"

func TestSnapshot() {

	caretaker := &caretaker{
		momentoArray: make([]*momento, 1),
	}

	originator := &originator{
		state: "aa",
	}

	caretaker.addMomento(originator.createMomento())

	originator.setState("bb")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMomento(originator.createMomento())

	originator.setState("cc")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMomento(originator.createMomento())

	originator.setState("dd")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	originator.restoreMomento(caretaker.getMomento(1))
	fmt.Printf("Originator Current State: %s\n", originator.getState())

}
