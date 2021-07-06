package wrapper

import "fmt"

func TestWrapper() {

	pizza := &veggeMania{}

	tomato := &tomatoTopping{
		pizza: pizza,
	}

	cheese := &cheeseTopping{
		pizza: tomato,
	}

	fmt.Printf("pizza final size is %d", cheese.getPrice())
}
