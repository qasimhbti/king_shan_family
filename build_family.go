package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type familyTree struct {
	Root *Member
}

type Member struct {
	id             int
	parentID       int
	name           string
	gender         string
	spouse         string
	spouseGender   string
	spouseParentID int
	offspring      []*Member
}

func (f *familyTree) addMember(id, parentID int, name, gender, spouse, spouseGender string, spouseParentID int) error {

	if f.Root == nil {
		f.Root = &Member{id: id,
			parentID:       parentID,
			name:           name,
			gender:         gender,
			spouse:         spouse,
			spouseGender:   spouseGender,
			spouseParentID: spouseParentID,
		}
		return nil
	}

	if f.Root.parentID == parentID {
		fmt.Printf("%s and %s already added\n", name, spouse)
		return nil
	}

	return f.Root.addMember(id, parentID, name, gender, spouse, spouseGender, spouseParentID)
}

func (m *Member) addMember(id, parentID int, name, gender, spouse, spouseGender string, spouseParentID int) error {
	/*if m == nil {
		return errors.New("Cannot add member to the family - King Shan family does not exists")
	}*/

	for _, v := range m.offspring {
		if name == v.name {
			fmt.Printf("%s and %s already added\n", name, spouse)
			return nil
		}
	}

	m.offspring = append(m.offspring, &Member{id: id,
		parentID:       parentID,
		name:           name,
		gender:         gender,
		spouse:         spouse,
		spouseGender:   spouseGender,
		spouseParentID: spouseParentID,
	})
	return nil
}

type newMember struct {
	id             int
	parentID       int
	name           string
	gender         string
	spouse         string
	spouseGender   string
	spouseParentID int
}

var shanFamily = []newMember{
	{2, 1, "KINGSHAN", "M", "QUEENANGA", "F", 0},
	{3, 2, "ISH", "M", "", "", 0},
	{4, 2, "CHIT", "M", "AMBI", "F", 0},
	{5, 2, "VICH", "M", "LIKA", "F", 0},
	{6, 2, "SATYA", "F", "VYAN", "M", 0},
	{7, 4, "DRITA", "M", "JAYA", "F", 0},
	{8, 4, "VRITA", "M", "", "", 0},
	{9, 5, "VILA", "M", "JNKI", "F", 0},
	{10, 5, "CHIKA", "F", "KPILA", "M", 0},
	{11, 6, "SATVY", "F", "ASVA", "M", 0},
	{12, 6, "SAVYA", "M", "KRPI", "F", 0},
	{13, 6, "SAAYAN", "M", "MINA", "F", 0},
	{14, 7, "JATA", "M", "", "", 0},
	{15, 7, "DRIYA", "F", "MNU", "M", 0},
	{16, 9, "LAVANYA", "F", "GRU", "M", 0},
	{17, 12, "KRIYA", "M", "", "", 0},
	{18, 13, "MISA", "M", "", "", 0},
}

func (k *familyTree) buildFamily(shanFamily []newMember) error {

	for _, m := range shanFamily {
		err := k.addMember(m.id, m.parentID, m.name, m.gender, m.spouse, m.spouseGender, m.spouseParentID)
		if err != nil {
			return errors.WithMessage(err, "buildFamily")
		}
	}

	return nil
}
