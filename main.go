package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		handleError(err)
	}
}

func run() error {
	kingShanFamily := &familyTree{}
	err := kingShanFamily.buildFamily(shanFamily)
	if err != nil {
		return errors.WithMessage(err, "build King Shan Family")
	}

	fmt.Println("****************************************************************")
	fmt.Println("Hello Geektrust, Welcome to the King Shan & Queen Anga Family!!")
	fmt.Println("****************************************************************")
	fmt.Println("Select the below options:-")
	fmt.Println("1. Meet the Family!")
	fmt.Println("2. A New Born!")
	fmt.Println("3. The Girl Child!")
	fmt.Println("4. Who's Your Daddy!")
	fmt.Println("Enter your option 1 or 2 or 3 or 4!")
	var option int
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		err := kingShanFamily.meetTheFamily()
		if err != nil {
			handleError(err)
		}
	case 2:
		err := kingShanFamily.newBorn()
		if err != nil {
			handleError(err)
		}
	case 3:
		err := kingShanFamily.girlChild()
		if err != nil {
			handleError(err)
		}
	case 4:
		/*err := kingShanFamily.yourDaddy()
		if err != nil {
			handleError(err)
		}*/
		fmt.Println("Not Implemented Yet!!")
	default:
		fmt.Println("Wrong Value Entered -- BYE")
	}

	return nil
}

func handleError(err error) {
	fmt.Println(err)
}
