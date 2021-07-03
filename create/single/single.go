package single

import (
	"fmt"
	"sync"
)

type single struct {

}

var lock = &sync.Mutex{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			println("begin create instance")
			singleInstance = &single{}
		} else {
			fmt.Println("instance has been created")
		}
	} else {
		fmt.Println("instance has been created")
	}

	return singleInstance
}


