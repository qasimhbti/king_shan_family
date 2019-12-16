package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (k *familyTree) newBorn() error {
	var mother, childName, childGender, childType string
	fmt.Println("Enter the Mother Name who is expecting a New Born Child")
	fmt.Scanf("%s", &mother)
	fmt.Println("Enter the Expected Child Name")
	fmt.Scanf("%s", &childName)
	fmt.Println("Whether the Expected Child is Son or Daughter --- Enter either son or daughter")
	fmt.Scanf("%s", &childType)

	mother = strings.ToUpper(mother)
	childName = strings.ToUpper(childName)
	childType = strings.ToUpper(childType)

	switch childType {
	case "SON":
		childGender = "M"
	case "DAUGHTER":
		childGender = "F"
	default:
		return errors.New("Please enter either Son or Daughter -- BYE!!!")
	}

	motherDetails, err := k.personDetails(mother)
	if err != nil {
		return errors.WithMessage(err, "NewBorn")
	}

	id := len(k.Root.offspring) + 3 // 2+1 - King Parent + Root ID

	c := &newMember{
		id:             id,
		parentID:       motherDetails.id,
		name:           childName,
		gender:         childGender,
		spouse:         "",
		spouseGender:   "",
		spouseParentID: 0,
	}

	err = k.addMember(c.id, c.parentID, c.name, c.gender, c.spouse, c.spouseGender, c.spouseParentID)
	if err != nil {
		return errors.WithMessage(err, "NewBorn")
	}
	fmt.Println("****************************************************************")
	fmt.Println("*** New Born Child Successfully added in King Shan Family ***")
	fmt.Println("****************************************************************")
	fmt.Println("*** Enter Person Name and Relation as follows ***")

	err = k.meetTheFamily()
	if err != nil {
		return errors.WithMessage(err, "NewBorn")
	}

	return nil
}
