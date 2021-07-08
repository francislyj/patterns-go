package intermediary

import "fmt"

type freightTrain struct {
	mediator mediator
}

func (f *freightTrain) arrive()  {
	if !f.mediator.canArrive(f) {
		fmt.Println("freight train: arrive blocked, waiting")
		return
	}

	fmt.Println("freight train arrived")
}

func (f *freightTrain) depart()  {
	fmt.Println("freight train : depart")
	f.mediator.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival()  {
	fmt.Println("freight train : arrival permitted")
	f.arrive()

}
