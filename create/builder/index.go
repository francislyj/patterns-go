package builder

import "fmt"

func TestBuilder() {
	normalBuilder := &normalBuilder{}

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("normal house window type is %s\n", normalHouse.windowType)
	fmt.Printf("normal house door type is %s\n", normalHouse.doorType)
	fmt.Printf("normal house floor is %d\n", normalHouse.floor)
}
