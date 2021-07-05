package adapter

import "fmt"

type windows struct {

}

func (w *windows) insertUsb() {
	fmt.Println("insert usb")
}
