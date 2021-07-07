package command

import "fmt"

type tv struct {

}

func (t *tv) on() {
	fmt.Println("tv on")
}

func (t *tv) off() {
	fmt.Println("tv off")
}
