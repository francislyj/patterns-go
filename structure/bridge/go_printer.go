package bridge

import "fmt"

type hpPrinter struct {

}

func (h *hpPrinter) printFile()  {
	fmt.Println("hp print file")
}