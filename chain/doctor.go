package chain

import "fmt"

type Doctor struct {
	Next department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.Next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.Next.Execute(p)
}

func (d *Doctor) SetNext(next department) {
	d.Next = next
}
