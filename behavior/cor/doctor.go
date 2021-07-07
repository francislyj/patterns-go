package cor

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("patient has already checked up")
		d.next.execute(p)
		return
	}

	fmt.Println("patient begin check up")
	p.doctorCheckUpDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}

func (d *doctor) setNext(next department)  {
	d.next = next
}
