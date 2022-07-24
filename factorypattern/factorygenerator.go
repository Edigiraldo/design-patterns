package factorypattern

import "fmt"

type Employee struct {
	Name         string `json:"name"`
	Title        string `json:"title"`
	AnnualSalary int    `json:"annual_salary"`
}

func (e Employee) Greet() {
	fmt.Printf("Hello, my name is %s and I'm a %s earning %d a year\n", e.Name, e.Title, e.AnnualSalary)
}

func NewEmployeeFactory(title string, annualSalary int) func(name string) Employee {
	return func(name string) Employee {
		return Employee{
			Name:         name,
			Title:        title,
			AnnualSalary: annualSalary,
		}
	}
}

func RunFactoryGeneratorExample() {
	frontendFactory := NewEmployeeFactory("Frontend Developer", 128000)
	Jane := frontendFactory("Jane")
	Andres := frontendFactory("Andres")

	BackendFactory := NewEmployeeFactory("Backend Developer", 128000)
	Daniel := BackendFactory("Daniel")
	Andrea := BackendFactory("Andrea")

	Andrea.Greet()
	Jane.Greet()
	Andres.Greet()
	Daniel.Greet()
}
