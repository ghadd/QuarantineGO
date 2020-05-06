package fhb

import "fmt"

// Whale class represents Whales
type Whale struct {
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

// LiveAnotherDay symbolized one Whales living day
func (p *Whale) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Whale) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another Whale
func (p *Whale) Reproduce() *Whale {
	return new(Whale)
}

// Die symbolizes Whales death
func (p *Whale) Die() {
	p.isAlive = false
}

// IsAlive returns fact of Whale's living
func (p *Whale) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Whale) Sound() {
	fmt.Println("Whale says hi wooooooohooooo!")
}

// Swim ...
func (p *Whale) Swim() {

}

// Display Displays info about a parrot
func (p *Whale) Display() {
	fmt.Println(*p)
}

func (w *Whale) Name() string {
	return w.name
}
