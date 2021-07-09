package observer

import "fmt"

type item struct {
	observerList []observer
	name string
	inStock bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("item update availability name %s \n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

func removeFromSlice(observerList []observer, removeOb observer) []observer {
	obLength := len(observerList)

	for i, observer := range observerList {
		if observer.getID() == removeOb.getID() {
			observerList[obLength - 1], observerList[i] = observerList[i], observerList[obLength - 1]
			return observerList[: obLength -1]
		}
	}

	return observerList
}