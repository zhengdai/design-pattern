package chain

import "fmt"

type Cachier struct {
	Next department
}

func (c *Cachier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cachier getting money from patient")
}

func (c *Cachier) SetNext(next department) {
	c.Next = next
}
