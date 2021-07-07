package cor

import "fmt"

type reception struct {
	next department
}

func (r *reception) setNext(next department)  {
	r.next = next
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("patient has already registration")
		r.next.execute(p)
		return
	}
	fmt.Println("patient begin registration")
	p.registrationDone = true
	if r.next != nil {
		r.next.execute(p)
	}
}
