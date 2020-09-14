package main

import (
	"bufio"
	"fmt"
	"os"
)

const appointedDay string = "30 october 1981"

type Citizen struct {
	isCitizen    bool
	IsAbandoned  bool
	IsBornUK     bool
	IsMinor      bool
	DateOfBirth  string
	PlaceOfBirth string
	IsResident   bool
	Father       *Citizen
	Mother       *Citizen
}

func NewCitizen(isAbandonned bool, isBornUK bool, isMinor bool, dateOfBirth string, placeOfBirth string, residency bool, c1 *Citizen, c2 *Citizen) *Citizen {
	return &Citizen{IsAbandoned: isAbandonned, IsBornUK: isBornUK, IsMinor: isMinor, DateOfBirth: dateOfBirth, PlaceOfBirth: placeOfBirth, IsResident: residency, Father: c1, Mother: c2}
}

func areParentResident(parent string) bool {
	fmt.Printf("Is your %s settled in the United Kingdom or that territory: (yes/no)", parent)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func areParentMilitary(parent string) bool {
	fmt.Printf("Is your %s member of armed forces: (yes/no)", parent)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func areParentCitizen(parent string) bool {
	fmt.Printf("Is your %s citizen: (yes/no)", parent)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func bornBeforeAct() bool {
	fmt.Printf("Are you born before the %s : (yes/no)", appointedDay)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func (Citizen) setDateOfBirth() string {
	fmt.Printf("When are you born ? ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func wasAbandonned() bool {
	fmt.Printf("Were you abandonned at birth : (yes/no)")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func main() {

	citizen := NewCitizen(false, false, false, "", "", false, nil, nil)
	citizen.Mother = NewCitizen(false, false, false, "", "", false, nil, nil)
	citizen.Father = NewCitizen(false, false, false, "", "", false, nil, nil)
	if bornBeforeAct() == true {
		fmt.Print("You are not a british citizen")
		return
	}
	citizen.IsAbandoned = wasAbandonned()
	if citizen.IsAbandoned == true {
		citizen.IsBornUK = askCountry(citizen)
		if citizen.IsBornUK == false {
			fmt.Print("You are not a british citizen")
		}
		citizen.Father.isCitizen = areParentCitizen("father")
		citizen.Mother.isCitizen = areParentCitizen("mother")
		citizen.Father.IsResident = areParentResident("father")
		citizen.Mother.IsResident = areParentResident("mother")
		if citizen.Father.isCitizen == true || citizen.Mother.isCitizen == true || citizen.Mother.IsResident || citizen.Father.IsResident {
			fmt.Print("You are a british citizen")
		} else {
			fmt.Print("You are not a british citizen")
		}
		return
	} else {
		citizen.isCitizen = areParentCitizen("mother") || areParentCitizen("father")
		if citizen.isCitizen {
			citizen.isCitizen = areParentResident("mother") || areParentResident("father")
		} else if citizen.isCitizen == false {
			citizen.isCitizen = areParentMilitary("mother") || areParentMilitary("father")
		}
	}
	citizen.IsMinor = askMinor()
	if citizen.IsMinor == true {
		citizen.Mother.IsResident = areParentResident("mother")
		citizen.Father.IsResident = areParentResident("father")
		citizen.Mother.isCitizen = areParentCitizen("mother")
		citizen.Father.isCitizen = areParentCitizen("father")
		citizen.isCitizen = citizen.Father.isCitizen == true || citizen.Mother.isCitizen == true || citizen.Mother.IsResident || citizen.Father.IsResident
	}
	if citizen.isCitizen == true {
		fmt.Print("You are a british citizen")
	} else {
		fmt.Print("You are not a british citizen")
	}

}

func askMinor() bool {
	fmt.Printf("Are you minor : (yes/no)")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text == "yes\n"
}

func askCountry(citizen *Citizen) bool {
	settlement := [3]string{"UK", "India", "Rep. Ireland"}
	fmt.Print("Country available :")
	fmt.Print(settlement)
	fmt.Printf("\nWhere where you abandonned : ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	var i int
	citizen.PlaceOfBirth = text
	for i = 0; i < len(settlement); i++ {
		if text == settlement[i] {
			return text == settlement[i]
		}
	}
	return text == settlement[i]
}
