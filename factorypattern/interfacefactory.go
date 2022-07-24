package factorypattern

import "fmt"

type Dog struct {
	name string
	age  int
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) SetAge(age int) {
	d.age = age
}

func (d Dog) Describe() {
	fmt.Printf("%s is a %d years old dog...\n", d.name, d.age)
}

type Duck struct {
	name string
	age  int
}

func (d *Duck) SetName(name string) {
	d.name = name
}

func (d *Duck) SetAge(age int) {
	d.age = age
}

func (d Duck) Describe() {
	fmt.Printf("%s is a %d years old duck...\n", d.name, d.age)
}

type IAnimal interface {
	SetName(name string)
	SetAge(age int)
	Describe()
}

type animalT string

var dogType animalT = "Dog"
var duckType animalT = "Duck"

func newAnimal(animalType animalT) (animal IAnimal) {
	switch animalType {
	case dogType:
		animal = &Dog{}
	case duckType:
		animal = &Duck{}
	default:
		fmt.Printf("%s type is not one of the expected ones...\n", animalType)
	}

	return
}

func RunInterfaceFactoryExample() {
	Max := newAnimal(dogType)
	Max.SetName("Max")
	Max.SetAge(2)

	Donald := newAnimal(duckType)
	Donald.SetName("Donald")
	Donald.SetAge(4)

	Max.Describe()
	Donald.Describe()
}
