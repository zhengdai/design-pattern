package test

import (
	"design-pattern/chain"
	"testing"
)

func TestChain(t *testing.T) {
	cachier := &chain.Cachier{}
	medical := &chain.Medical{}
	medical.SetNext(cachier)
	doctor := &chain.Doctor{}
	doctor.SetNext(medical)
	reception := &chain.Reception{}
	reception.SetNext(doctor)
	patient := &chain.Patient{Name: "abc"}
	reception.Execute(patient)
}
