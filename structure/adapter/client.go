package adapter

import "fmt"

type client struct {
	
}

func (c *client) insertIntoLightningPortToComputer(com computer) {
	fmt.Println("client insert into lightning port to computer")
	com.insertIntoLightningPort()
}
