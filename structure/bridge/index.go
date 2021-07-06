package bridge

func TestBridge() {
	hp := &hpPrinter{}
	epson := &epson{}

	windows := &windows{}

	windows.setPrinter(hp)
	windows.printFile()
	windows.setPrinter(epson)
	windows.printFile()

	mac := &mac{}
	mac.setPrinter(hp)
	mac.printFile()
	mac.setPrinter(epson)
	mac.printFile()
}
