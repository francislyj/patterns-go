package bridge

type windows struct {
	printer
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

func (w *windows) printFile() {
	w.printer.printFile()
}
