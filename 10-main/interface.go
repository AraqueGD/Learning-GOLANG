package main

import "fmt"

type lifes interface {
	IsLife() bool
}
type human interface {
	breathe()
	think()
	eat()
	sex() string
	personal() int
}

type animal interface {
	breathe()
	eat()
	carnivorous() bool
}

type vegetal interface {
	ClasificationVegetal() string
}

// Genre Human
type man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	IsMan     bool
	life      bool
}

type woman struct {
	man
}

func (m *man) breathe()      { m.breathing = true }
func (m *man) eat()          { m.eating = true }
func (m *man) think()        { m.thinking = true }
func (m *man) personal() int { return m.age }
func (m *man) sex() string {
	if m.IsMan == true {
		return "Man"
	} else {
		return "Woman"
	}
}
func (m *man) IsLife() bool { return m.life }
func Humanbreathing(hu human) {
	hu.breathe()
	fmt.Printf("I´am %s and already breathing\n", hu.sex())
}

func PersonalInformation(hu human) {
	fmt.Printf("My Personal Information is: Age: %d\n", hu.personal())
}

// Dogs
type Dog struct {
	breathing  bool
	eating     bool
	carnivorou bool
	life       bool
}

func (p *Dog) breathe()          { p.breathing = true }
func (p *Dog) eat()              { p.eating = true }
func (p *Dog) carnivorous() bool { return p.carnivorou }
func (p *Dog) IsLife() bool      { return p.life }

func Animalbreathing(an animal) {
	an.breathe()
	fmt.Printf("I´am Animal and already breathing\n")
}

func AnimalCarnivous(an animal) int {
	if an.carnivorous() == true {
		return 1
	}
	return 0
}
func IamLife(l lifes) bool {
	return l.IsLife()
}

func main() {
	Camilo := new(man)
	Camilo.IsMan = true
	Humanbreathing(Camilo)

	Maria := new(woman)
	Humanbreathing(Maria)

	totalCarnivorous := 0
	milo := new(Dog)
	milo.carnivorou = false
	wanda := new(Dog)
	wanda.carnivorou = true
	Animalbreathing(wanda)
	Animalbreathing(milo)
	totalCarnivorous += AnimalCarnivous(milo)
	totalCarnivorous += AnimalCarnivous(wanda)

	fmt.Printf("Total Carnivorous %d\n", totalCarnivorous)

	Camilo.age = 15

	PersonalInformation(Camilo)
	milo.life = true
	Camilo.life = true
	fmt.Printf("I am Life = %t\n", IamLife(Camilo))
	fmt.Printf("I am Life = %t", IamLife(milo))

}
