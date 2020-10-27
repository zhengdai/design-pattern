package mediator

import "fmt"

type PassengerTrain struct {
	Mediator mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("passengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}
