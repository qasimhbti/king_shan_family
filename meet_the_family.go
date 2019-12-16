package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (k *familyTree) meetTheFamily() error {
	var person, relation string
	fmt.Println("Enter Person Name")
	fmt.Scanf("%s", &person)
	fmt.Println("Enter the Relation")
	fmt.Scanf("%s", &relation)

	person = strings.ToUpper(person)
	relation = strings.ToUpper(relation)

	switch relation {
	case "BROTHERS", "BROTHER":
		brothers, err := k.relationBrothers(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(brothers, person, relation)
	case "SISTERS", "SISTER":
		sisters, err := k.relationSisters(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(sisters, person, relation)
	case "FATHER":
		father, err := k.relationFather(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(father, person, relation)
	case "MOTHER":
		mother, err := k.relationMother(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(mother, person, relation)
	case "CHILDREN":
		children, err := k.relationChildren(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(children, person, relation)
	case "SON", "SONS":
		son, err := k.relationSon(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(son, person, relation)
	case "DAUGHTER", "DAUGHTERS":
		daughter, err := k.relationDaughter(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(daughter, person, relation)
	case "COUSINS":
		cousins, err := k.relationCousins(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(cousins, person, relation)
	case "GRANDDAUGHTER":
		grandDaughters, err := k.relationGrandDaughter(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(grandDaughters, person, relation)
	case "GRANDSON":
		grandSons, err := k.relationGrandSon(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(grandSons, person, relation)
	case "GRANDCHILDREN":
		grandChildren, err := k.relationGrandChildren(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(grandChildren, person, relation)
	case "BROTHERINLAW":
		brotherInLaw, err := k.relationBrotherInLaw(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(brotherInLaw, person, relation)
	case "SISTERINLAW":
		sisterInLaw, err := k.relationSisterInLaw(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(sisterInLaw, person, relation)
	case "PATERNALAUNT":
		paternalAunts, err := k.relationPaternalAunts(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(paternalAunts, person, relation)
	case "MATERNALAUNT":
		maternalAunts, err := k.relationMaternalAunts(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(maternalAunts, person, relation)
	case "PATERNALUNCLE":
		paternalUncles, err := k.relationPaternalUncles(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(paternalUncles, person, relation)
	case "MATERNALUNCLE":
		maternalUncles, err := k.relationMaternalUncles(person)
		if err != nil {
			return errors.WithMessage(err, "meet_the_family")
		}
		logRelation(maternalUncles, person, relation)
	default:
		return errors.Errorf("%s relation does not exists", relation)
	}
	return nil
}

func (k *familyTree) personDetails(person string) (*newMember, error) {
	member := &newMember{}

	if person == k.Root.name {
		member.id = k.Root.id
		member.parentID = k.Root.parentID
		member.name = k.Root.name
		member.gender = k.Root.gender
		member.spouse = k.Root.spouse
		member.spouseGender = k.Root.spouseGender
		member.spouseParentID = k.Root.spouseParentID
		return member, nil
	} else if person == k.Root.spouse {
		member.id = k.Root.id
		member.parentID = k.Root.spouseParentID
		member.name = k.Root.spouse
		member.gender = k.Root.spouseGender
		member.spouse = k.Root.name
		member.spouseGender = k.Root.gender
		member.spouseParentID = k.Root.parentID
		return member, nil
	}

	for _, m := range k.Root.offspring {
		if person == m.name {
			member.id = m.id
			member.parentID = m.parentID
			member.name = m.name
			member.gender = m.gender
			member.spouse = m.spouse
			member.spouseGender = m.spouseGender
			member.spouseParentID = m.spouseParentID
		} else if person == m.spouse {
			member.id = m.id
			member.parentID = m.spouseParentID
			member.name = m.spouse
			member.gender = m.spouseGender
			member.spouse = m.name
			member.spouseGender = m.gender
			member.spouseParentID = m.parentID
		}
	}

	if member.name == "" {
		return member, errors.Errorf("%s does not belong to King Shan Family", person)
	}

	return member, nil

}

func (k *familyTree) relationBrothers(person string) ([]string, error) {
	var brothers []string

	member, err := k.personDetails(person)
	if err != nil {
		return brothers, errors.WithMessage(err, "relationBrothers")
	}

	for _, m := range k.Root.offspring {
		if member.parentID == m.parentID && m.gender == "M" && member.name != m.name {
			brothers = append(brothers, m.name)
		}
	}

	return brothers, nil
}

func (k *familyTree) relationSisters(person string) ([]string, error) {
	var sisters []string

	member, err := k.personDetails(person)
	if err != nil {
		return sisters, errors.WithMessage(err, "relationSisters")
	}

	for _, m := range k.Root.offspring {
		if member.parentID == m.parentID && m.gender == "F" && member.name != m.name {
			sisters = append(sisters, m.name)
		}
	}

	return sisters, nil
}

func (k *familyTree) relationFather(person string) ([]string, error) {
	var father []string

	member, err := k.personDetails(person)
	if err != nil {
		return father, errors.WithMessage(err, "relationFather")
	}

	if member.parentID == k.Root.id && k.Root.gender == "M" {
		father = append(father, k.Root.name)
		return father, nil
	}

	for _, m := range k.Root.offspring {
		if member.parentID == m.id && m.gender == "M" {
			father = append(father, m.name)
		}
	}

	return father, nil
}

func (k *familyTree) relationMother(person string) ([]string, error) {
	var mother []string

	member, err := k.personDetails(person)
	if err != nil {
		return mother, errors.WithMessage(err, "relationMother")
	}

	if member.parentID == k.Root.id && k.Root.spouseGender == "F" {
		mother = append(mother, k.Root.spouse)
		return mother, nil
	}

	for _, m := range k.Root.offspring {
		if member.parentID == m.id && m.spouseGender == "F" {
			mother = append(mother, m.spouse)
		}
	}

	return mother, nil
}

func (k *familyTree) relationChildren(person string) ([]string, error) {
	var children []string

	member, err := k.personDetails(person)
	if err != nil {
		return children, errors.WithMessage(err, "relationChildren")
	}

	for _, m := range k.Root.offspring {
		if member.id == m.parentID {
			children = append(children, m.name)
		}
	}

	return children, nil
}

func (k *familyTree) relationSon(person string) ([]string, error) {
	var son []string

	member, err := k.personDetails(person)
	if err != nil {
		return son, errors.WithMessage(err, "relationSon")
	}

	for _, m := range k.Root.offspring {
		if member.id == m.parentID && m.gender == "M" {
			son = append(son, m.name)
		}
	}

	return son, nil
}

func (k *familyTree) relationDaughter(person string) ([]string, error) {
	var daughter []string

	member, err := k.personDetails(person)
	if err != nil {
		return daughter, errors.WithMessage(err, "relationDaughter")
	}

	for _, m := range k.Root.offspring {
		if member.id == m.parentID && m.gender == "F" {
			daughter = append(daughter, m.name)
		}
	}

	return daughter, nil
}

func (k *familyTree) relationCousins(person string) ([]string, error) {
	var parentsBroSis, cousins []string

	father, err := k.relationFather(person)
	if err != nil {
		return cousins, errors.WithMessage(err, "relationCousins")
	}

	mother, err := k.relationMother(person)
	if err != nil {
		return cousins, errors.WithMessage(err, "relationCousins")
	}

	parents := append([]string{}, append(father, mother...)...)

	for _, p := range parents {
		brothers, err := k.relationBrothers(p)
		if err != nil {
			return cousins, errors.WithMessage(err, "relationCousins")
		}
		sisters, err := k.relationSisters(p)
		if err != nil {
			return cousins, errors.WithMessage(err, "relationCousins")
		}

		parentsBroSis = append(parentsBroSis, append(brothers, sisters...)...)
	}

	for _, p := range parentsBroSis {
		children, err := k.relationChildren(p)
		if err != nil {
			return cousins, errors.WithMessage(err, "relationCousins")
		}

		cousins = append([]string{}, append(cousins, children...)...)
	}

	uCousins := removeDuplicate(cousins)

	return uCousins, nil
}

func (k *familyTree) relationGrandDaughter(person string) ([]string, error) {
	var grandDaughter []string

	children, err := k.relationChildren(person)
	if err != nil {
		return grandDaughter, errors.WithMessage(err, "relationGrandDaughter")
	}

	for _, c := range children {
		daughter, err := k.relationDaughter(c)
		if err != nil {
			return grandDaughter, errors.WithMessage(err, "relationGrandDaughter")
		}

		grandDaughter = append([]string{}, append(grandDaughter, daughter...)...)
	}

	return grandDaughter, nil
}

func (k *familyTree) relationGrandSon(person string) ([]string, error) {
	var grandSons []string

	children, err := k.relationChildren(person)
	if err != nil {
		return grandSons, errors.WithMessage(err, "relationGrandSon")
	}

	for _, c := range children {
		sons, err := k.relationSon(c)
		if err != nil {
			return grandSons, errors.WithMessage(err, "relationGrandSon")
		}

		grandSons = append([]string{}, append(grandSons, sons...)...)
	}

	return grandSons, nil
}

func (k *familyTree) relationGrandChildren(person string) ([]string, error) {
	var grandChildren []string

	children, err := k.relationChildren(person)
	if err != nil {
		return grandChildren, errors.WithMessage(err, "relationGrandChildren")
	}

	for _, c := range children {
		children, err := k.relationChildren(c)
		if err != nil {
			return grandChildren, errors.WithMessage(err, "relationGrandChildren")
		}

		grandChildren = append([]string{}, append(grandChildren, children...)...)
	}

	return grandChildren, nil
}

func (k *familyTree) relationBrotherInLaw(person string) ([]string, error) {
	var brotherInLaw []string

	member, err := k.personDetails(person)
	if err != nil {
		return brotherInLaw, errors.WithMessage(err, "relationBrotherInLaw")
	}

	brothers, err := k.relationBrothers(member.spouse)
	if err != nil {
		return brotherInLaw, errors.WithMessage(err, "relationBrotherInLaw")
	}

	brotherInLaw = append([]string{}, append(brotherInLaw, brothers...)...)

	siblings, err := k.relationSiblings(member.name)
	if err != nil {
		return brotherInLaw, errors.WithMessage(err, "relationBrotherInLaw")
	}

	for _, s := range siblings {
		member, err := k.personDetails(s)
		if err != nil {
			return brotherInLaw, errors.WithMessage(err, "relationBrotherInLaw")
		}

		if member.gender == "F" && member.spouse != "" {
			brotherInLaw = append(brotherInLaw, member.spouse)
		}
	}

	return brotherInLaw, nil
}

func (k *familyTree) relationSisterInLaw(person string) ([]string, error) {
	var sisterInLaw []string

	member, err := k.personDetails(person)
	if err != nil {
		return sisterInLaw, errors.WithMessage(err, "relationSisterInLaw")
	}

	sisters, err := k.relationSisters(member.spouse)
	if err != nil {
		return sisterInLaw, errors.WithMessage(err, "relationSisterInLaw")
	}

	sisterInLaw = append([]string{}, append(sisterInLaw, sisters...)...)

	siblings, err := k.relationSiblings(member.name)
	if err != nil {
		return sisterInLaw, errors.WithMessage(err, "relationSisterInLaw")
	}

	for _, s := range siblings {
		member, err := k.personDetails(s)
		if err != nil {
			return sisterInLaw, errors.WithMessage(err, "relationSisterInLaw")
		}

		if member.gender == "M" && member.spouse != "" {
			sisterInLaw = append(sisterInLaw, member.spouse)
		}
	}

	return sisterInLaw, nil
}

func (k *familyTree) relationPaternalAunts(person string) ([]string, error) {
	var paternalAunts []string

	father, err := k.relationFather(person)
	if err != nil {
		return paternalAunts, errors.WithMessage(err, "relationPaternalAunts")
	}

	for _, f := range father {
		sisters, err := k.relationSisters(f)
		if err != nil {
			return paternalAunts, errors.WithMessage(err, "relationPaternalAunts")
		}

		paternalAunts = append([]string{}, append(paternalAunts, sisters...)...)

		sistersInLaw, err := k.relationSisterInLaw(f)
		if err != nil {
			return paternalAunts, errors.WithMessage(err, "relationPaternalAunts")
		}

		paternalAunts = append([]string{}, append(paternalAunts, sistersInLaw...)...)
	}

	return paternalAunts, nil
}

func (k *familyTree) relationMaternalAunts(person string) ([]string, error) {
	var maternalAunts []string

	mother, err := k.relationMother(person)
	if err != nil {
		return maternalAunts, errors.WithMessage(err, "relationMaternalAunts")
	}

	for _, m := range mother {
		sisters, err := k.relationSisters(m)
		if err != nil {
			return maternalAunts, errors.WithMessage(err, "relationMaternalAunts")
		}

		maternalAunts = append([]string{}, append(maternalAunts, sisters...)...)

		sistersInLaw, err := k.relationSisterInLaw(m)
		if err != nil {
			return maternalAunts, errors.WithMessage(err, "relationMaternalAunts")
		}

		maternalAunts = append([]string{}, append(maternalAunts, sistersInLaw...)...)
	}

	return maternalAunts, nil
}

func (k *familyTree) relationPaternalUncles(person string) ([]string, error) {
	var paternalUncles []string

	father, err := k.relationFather(person)
	if err != nil {
		return paternalUncles, errors.WithMessage(err, "relationPaternalUncles")
	}

	for _, f := range father {
		brothers, err := k.relationBrothers(f)
		if err != nil {
			return paternalUncles, errors.WithMessage(err, "relationPaternalUncles")
		}

		paternalUncles = append([]string{}, append(paternalUncles, brothers...)...)

		brothersInLaw, err := k.relationBrotherInLaw(f)
		if err != nil {
			return paternalUncles, errors.WithMessage(err, "relationPaternalUncles")
		}

		paternalUncles = append([]string{}, append(paternalUncles, brothersInLaw...)...)
	}

	return paternalUncles, nil
}

func (k *familyTree) relationMaternalUncles(person string) ([]string, error) {
	var maternalUncles []string

	mother, err := k.relationMother(person)
	if err != nil {
		return maternalUncles, errors.WithMessage(err, "relationMaternalUncles")
	}

	for _, m := range mother {
		brothers, err := k.relationBrothers(m)
		if err != nil {
			return maternalUncles, errors.WithMessage(err, "relationMaternalUncles")
		}

		maternalUncles = append([]string{}, append(maternalUncles, brothers...)...)

		brothersInLaw, err := k.relationBrotherInLaw(m)
		if err != nil {
			return maternalUncles, errors.WithMessage(err, "relationMaternalUncles")
		}

		maternalUncles = append([]string{}, append(maternalUncles, brothersInLaw...)...)
	}

	return maternalUncles, nil
}

func (k *familyTree) relationSiblings(person string) ([]string, error) {
	var siblings []string

	brothers, err := k.relationBrothers(person)
	if err != nil {
		return siblings, errors.WithMessage(err, "relationSiblings")
	}
	siblings = append([]string{}, append(siblings, brothers...)...)

	sisters, err := k.relationSisters(person)
	if err != nil {
		return siblings, errors.WithMessage(err, "relationSiblings")
	}
	siblings = append([]string{}, append(siblings, sisters...)...)

	return siblings, nil
}

func logRelation(member []string, person, relation string) {
	if len(member) == 0 {
		fmt.Println("****************************************************************")
		fmt.Printf("%s has no %s.\n", person, relation)
		fmt.Println("****************************************************************")
		return
	}

	fmt.Println("****************************************************************")
	fmt.Printf("%s of %s is/are as follows::\n", relation, person)
	fmt.Println("****************************************************************")
	for _, m := range member {
		fmt.Println(m)
	}
}

func removeDuplicate(items []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range items {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
