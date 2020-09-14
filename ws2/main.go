package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Rule func(*Citizen)bool

type Parent struct {
	father *Citizen
	mother *Citizen
}

type Citizen struct {
	bornInUK              bool
	bornAfterCommencement bool
	bornInQualifyingTerritory bool
	isSettledInQualifyingTerritory bool
	bornAfterAppointedDay bool
	parents               *Parent
	isBritish             bool
}

func (c *Citizen) BornAfterAppointedDay(citizen *Citizen) bool {
	return c.bornAfterAppointedDay
}

func (c *Citizen) SetBornAfterAppointedDay(bornAfterAppointedDay bool) {
	c.bornAfterAppointedDay = bornAfterAppointedDay
}

func (c *Citizen) IsSettledInQualifyingTerritory(citizen *Citizen) bool {
	return c.isSettledInQualifyingTerritory
}

func (c *Citizen) SetIsSettledInQualifyingTerritory(isSettledInQualifyingTerritory bool) {
	c.isSettledInQualifyingTerritory = isSettledInQualifyingTerritory
}

func (c *Citizen) BornInQualifyingTerritory(citizen *Citizen) bool {
	return c.bornInQualifyingTerritory
}

func (c *Citizen) SetBornInQualifyingTerritory(bornInQualifyingTerritory bool) {
	c.bornInQualifyingTerritory = bornInQualifyingTerritory
}

func NewParent(father *Citizen, mother *Citizen) *Parent {
	return &Parent{father: father, mother: mother}
}

func NewCitizen(bornInUK bool, bornAfterCommencement bool, parents *Parent, isBritish bool) *Citizen {
	return &Citizen{bornInUK: bornInUK, bornAfterCommencement: bornAfterCommencement, parents: parents, isBritish: isBritish}
}

func (c *Citizen) IsBritish(citizen *Citizen) bool {
	return c.isBritish
}

func (c *Citizen) SetIsBritish(isBritish bool) {
	c.isBritish = isBritish
}

func (c *Citizen) BornAfterCommencement(citizen *Citizen) bool {
	return c.bornAfterCommencement
}

func (c *Citizen) SetBornAfterCommencement(bornAfterCommencement bool) {
	c.bornAfterCommencement = bornAfterCommencement
}

func (c *Citizen) BornInUK(citizen *Citizen) bool {
	return c.bornInUK
}

func (c *Citizen) SetBornInUK(bornInUK bool) {
	c.bornInUK = bornInUK
}

func (c *Citizen) Parents() *Parent {
	return c.parents
}

func (c *Citizen) SetParents(parents *Parent) {
	c.parents = parents
}

func (p *Parent) parentIsCitizen(parent *Citizen) bool {
	return isCitizen(parent)
}

func isCitizen(citizen *Citizen) bool {

	if citizen.bornAfterCommencement && citizen.bornInUK {
		citizen.isBritish = true
	}
	if citizen.Parents() != nil {
		if citizen.Parents().parentIsCitizen(citizen.Parents().father) || citizen.Parents().parentIsCitizen(citizen.Parents().mother) {
			citizen.isBritish = true
		}
	}
	return citizen.isBritish
}
func isSettledInQualifyingTerritory(citizen *Citizen)bool {
	return citizen.isBritish
}

func main() {
	citizen := NewCitizen(rand.Float32() > 0.5, rand.Float32() > 0.5, NewParent(NewCitizen(rand.Float32() > 0.5, rand.Float32() > 0.5, nil, rand.Float32() > 0.5), NewCitizen(rand.Float32() > 0.5, rand.Float32() > 0.5, nil, rand.Float32() > 0.5)), false)

	RulesMap := map[string]Rule{
		"britishCitizen":isCitizen,
		"isBritish":citizen.IsBritish,
		"isSettledInQualifyingTerritory":citizen.IsSettledInQualifyingTerritory,
		"bornInQualifyingTerritory":citizen.BornInQualifyingTerritory,
		"bornInUk":citizen.BornInUK,
		"bornAfterCommencement":citizen.BornAfterCommencement,
		"bornAfterAppointedDay":citizen.BornAfterAppointedDay,
	}


	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(citizen)
	fmt.Print(isCitizen(citizen))

}
