package bridge

type mac struct {
	printer
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

func (m *mac) printFile() {
	m.printer.printFile()
}
