package parrot

import "fmt"

// Parrot class represents parrots
type Parrot struct {
	isAlive bool

	name               string
	daysAlive          uint16
	sex                string
	age                uint8
	isReproductive     bool
	reproductionMinAge uint8

	hunger uint8
	health uint8
}

// LiveAnotherDay symbolized one parrots living day
func (p *Parrot) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Parrot) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another parrot
func (p *Parrot) Reproduce() *Parrot {
	return new(Parrot)
}

// Die symbolizes parrots death
func (p *Parrot) Die() {
	p.isAlive = false
}

// IsAlive returns fact of parrot's living
func (p *Parrot) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Parrot) Sound() {
	fmt.Println("ARR. I'm the parrot!")
}
