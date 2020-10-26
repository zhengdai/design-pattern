package chain

import "fmt"

type Reception struct {
	Next department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patinet registration already done")
		r.Next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next department) {
	r.Next = next
}
