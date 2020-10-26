package chain

import "fmt"

type Medical struct {
	Next department
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.Next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	m.Next.Execute(p)
}

func (m *Medical) SetNext(next department) {
	m.Next = next
}
