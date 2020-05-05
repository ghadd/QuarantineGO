package fpd

import "fmt"

// Shark class represents Sharks
type Shark struct {
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

// LiveAnotherDay symbolized one Sharks living day
func (p *Shark) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Shark) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another Shark
func (p *Shark) Reproduce() *Shark {
	return new(Shark)
}

// Die symbolizes Sharks death
func (p *Shark) Die() {
	p.isAlive = false
}

// IsAlive returns fact of Shark's living
func (p *Shark) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Shark) Sound() {
	fmt.Println("Shark says hi and bye!")
}

// Swim ...
func (p *Shark) Swim() {

}

// Display Displays info about a parrot
func (p *Shark) Display() {
	fmt.Println(*p)
}
