package single

import "fmt"

func TestSingle() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
