package main

import (
	"fmt"
	"github.com/francislyj/patterns-go/structure/bridge"
	"sync"
)

func TestMutex() {
	var mutex sync.Mutex
	fmt.Printf("%+v\n", mutex)

	mutex.Lock()
	fmt.Printf("%+v\n", mutex)

	mutex.Unlock()
	fmt.Printf("%+v\n", mutex)
}

func main() {
	//factory.TestFactory()
	//factory_abs.TestAbsFactory()
	//builder.TestBuilder()

	//prototype.TestPrototype()

	//single.TestSingle()


	//TestMutex()

	//adapter.TestAdapter()

	bridge.TestBridge()



}
