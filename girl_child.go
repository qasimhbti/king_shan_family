package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (k *familyTree) girlChild() error {
	mothers, err := k.motherWithMaxDaughters()
	if err != nil {
		return errors.WithMessage(err, "GirlChild")
	}

	fmt.Println("**** Following Mothers has the most Girl Children ****")
	for _, m := range mothers {
		fmt.Println(m)
	}

	err = k.addChild()
	if err != nil {
		return errors.WithMessage(err, "GirlChild")
	}
	fmt.Println("****************************************************************")
	fmt.Println("** New Born Child Successfully added in King Shan Family **")
	fmt.Println("****************************************************************")

	mothers, err = k.motherWithMaxDaughters()
	if err != nil {
		return errors.WithMessage(err, "GirlChild")
	}

	fmt.Println("**** Following Mothers has the most Girl Children ****")
	for _, m := range mothers {
		fmt.Println(m)
	}
	return nil
}

func (k *familyTree) motherWithMaxDaughters() ([]string, error) {
	parent := make(map[string]int)
	maxNumOfDaughters := 0
	var parentWithMaxNumOfDaughters, mothersWithMaxNumOfDaughters []string

	for _, m := range k.Root.offspring {

		daughters, err := k.relationDaughter(m.name)
		if err != nil {
			return mothersWithMaxNumOfDaughters, errors.WithMessage(err, "MotherWithMaxDaughters")
		}
		parent[m.name] = len(daughters)
	}

	for _, n := range parent {
		if n >= maxNumOfDaughters {
			maxNumOfDaughters = n
		}
	}

	for p, n := range parent {
		if n == maxNumOfDaughters {
			parentWithMaxNumOfDaughters = append(parentWithMaxNumOfDaughters, p)
		}
	}

	for _, p := range parentWithMaxNumOfDaughters {
		member, err := k.personDetails(p)
		if err != nil {
			return mothersWithMaxNumOfDaughters, errors.WithMessage(err, "MotherWithMaxDaughters")
		}

		if member.gender == "F" {
			mothersWithMaxNumOfDaughters = append(mothersWithMaxNumOfDaughters, member.name)
		} else {
			mothersWithMaxNumOfDaughters = append(mothersWithMaxNumOfDaughters, member.spouse)
		}
	}

	return mothersWithMaxNumOfDaughters, nil
}

func (k *familyTree) addChild() error {
	var mother, childName, childGender, childType string
	fmt.Println("**** Enter the following details for addition of new child ****")
	fmt.Println("Enter the Mother Name!!!")
	fmt.Scanf("%s", &mother)
	fmt.Println("Enter the Child Name")
	fmt.Scanf("%s", &childName)
	fmt.Println("Whether the Expected Child is Son or Daughter --- Enter son or daughter")
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
		return errors.New("Please enter either Son or Daughter")
	}

	motherDetails, err := k.personDetails(mother)
	if err != nil {
		return errors.WithMessage(err, "AddChild")
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
		return errors.WithMessage(err, "AddChild")
	}

	return nil
}
