package bhb

import "fmt"

// Parrot class represents parrots
type Parrot struct {
	isAlive bool

	name               string
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

// Reproduct returns a new instance os a class
// func Reproduct() *Parrot {
// 		return new(Parrot)
// }

// Display Displays info about a parrot
func (p *Parrot) Display() {
	fmt.Println(*p)
}

// NewParrot creates an instance with given params
func NewParrot(name, sex string, isRep bool, repMinAge uint8) *Parrot {
	return &Parrot{
		isAlive:            true,
		name:               name,
		sex:                sex,
		age:                0,
		isReproductive:     isRep,
		reproductionMinAge: repMinAge,
		hunger:             0,
		health:             0,
	}
}

// Fly ...
func (p *Parrot) Fly() {

}
