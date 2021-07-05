package adapter

import "fmt"

type mac struct {

}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("mac insert into lighting port")
}
