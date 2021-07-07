package cor

type department interface {
	execute(*patient)
	setNext(department)
}
