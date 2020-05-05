package bpd

import "fmt"

// Eagle class represents Eagles
type Eagle struct {
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

// LiveAnotherDay symbolized one Eagles living day
func (p *Eagle) LiveAnotherDay() {

}

// Eat eats a certain amount of food
func (p *Eagle) Eat(amount uint8) {
	p.hunger += amount
}

// Reproduce reproduces another Eagle
func (p *Eagle) Reproduce() *Eagle {
	return new(Eagle)
}

// Die symbolizes Eagles death
func (p *Eagle) Die() {
	p.isAlive = false
}

// IsAlive returns fact of Eagle's living
func (p *Eagle) IsAlive() bool {
	return p.isAlive
}

// Sound makes some sound
func (p *Eagle) Sound() {
	fmt.Println("Eagle goes brrrr!")
}

// Reproduct returns a new instance os a class
// func Reproduct() *Eagle {
// 	return new(Eagle)
// }

// Display Displays info about a eagle
func (p *Eagle) Display() {
	fmt.Println(*p)
}

// Fly ...
func (p *Eagle) Fly() {

}
