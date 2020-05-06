package mhb

import "fmt"

// Giraffe class represents Giraffes
type Giraffe struct {
	isAlive            bool
	name               string
	daysAlive          uint16
	sex                string
	age                uint8
	isReproductive     bool
	reproductionMinAge uint8

	hunger uint8
	health uint8
}

// LiveAnotherDay symbolized one Giraffes living day
func (p *Giraffe) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Giraffe) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another Giraffe
func (p *Giraffe) Reproduce() *Giraffe {
	return new(Giraffe)
}

// Die symbolizes Giraffes death
func (p *Giraffe) Die() {
	p.isAlive = false
}

// IsAlive returns fact of Giraffe's living
func (p *Giraffe) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Giraffe) Sound() {
	fmt.Println("Giraffe says hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii!")
}

// DoSex method symbolizes all mammal creatures
func (p *Giraffe) DoSex() {

}

// Display Displays info about a parrot
func (p *Giraffe) Display() {
	fmt.Println(*p)
}

func (g *Giraffe) Name() string {
	return g.name
}
