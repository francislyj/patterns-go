package cor

import "fmt"

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("patient has already medical done")
		m.next.execute(p)
		return
	}

	fmt.Println("patient begin medical")
	p.medicineDone = true
	if m.next != nil {
		m.next.execute(p)
	}
}

func (m *medical) setNext(next department) {
	m.next = next
}
