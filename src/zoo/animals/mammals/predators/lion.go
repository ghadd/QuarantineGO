package mpd

import "fmt"

// Lion class represents Lions
type Lion struct {
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

// LiveAnotherDay symbolized one Lions living day
func (p *Lion) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Lion) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another Lion
func (p *Lion) Reproduce() *Lion {
	return new(Lion)
}

// Die symbolizes Lions death
func (p *Lion) Die() {
	p.isAlive = false
}

// IsAlive returns fact of Lion's living
func (p *Lion) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Lion) Sound() {
	fmt.Println("Lion says aarrrrr!")
}

// DoSex method symbolizes all mammal creatures
func (p *Lion) DoSex() {

}

// Display Displays info about a parrot
func (p *Lion) Display() {
	fmt.Println(*p)
}
