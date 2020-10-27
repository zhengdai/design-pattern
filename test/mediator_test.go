package test

import (
	"design-pattern/mediator"
	"testing"
)

func TestMediator(t *testing.T) {
	stationManager := mediator.NewStationManager()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}

	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
