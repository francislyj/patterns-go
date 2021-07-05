package adapter

import "fmt"

type windowAdapter struct {
	windows *windows
}

func (w *windowAdapter) insertIntoLightningPort() {
	fmt.Println("window adapter insert into lightning port")
	w.windows.insertUsb()
}
