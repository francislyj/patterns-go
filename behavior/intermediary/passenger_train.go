package intermediary

import (
	"fmt"
)

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("passenger train blocked, waiting")
		return
	}

	fmt.Println("passenger train arrived")
}

func (p *passengerTrain) depart() {
	fmt.Println("passenger train : leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival()  {
	fmt.Println("passenger train : arrival permit, arriving")
	p.arrive()
}
