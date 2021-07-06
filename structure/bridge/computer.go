package bridge

type computer interface {
	setPrinter(printer)
	print()
}
