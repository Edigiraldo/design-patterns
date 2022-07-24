package factorypattern

import "fmt"

type Person struct {
	name   string
	age    int
	height float64
	email  string
}

func (p Person) Contact() {
	fmt.Printf("Writting to %s through %s...\n", p.name, p.email)
}

func newPerson(name string, age int, email string, height float64) Person {
	return Person{
		name:   name,
		age:    age,
		email:  email,
		height: height,
	}
}

func RunFactoryExample() {
	Andrea := newPerson("Andrea", 50, "andrea@email.com", 1.64)
	Rob := newPerson("Rob", 31, "rob@email.com", 1.75)

	Andrea.Contact()
	Rob.Contact()
}
