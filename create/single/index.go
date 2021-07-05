package single

import "fmt"

func TestSingle() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	var name string

	fmt.Scanln(&name)

	fmt.Printf("name is %s", name)
}
