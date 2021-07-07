package cor

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("patient has already payment")
		return
	}

	fmt.Println("patient begin to payment")
	p.paymentDone = true
	if c.next != nil {
		c.next.execute(p)
	}

}

func (c *cashier) setNext(next department) {
	c.next = next
}