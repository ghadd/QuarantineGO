package bhb

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newHealth := r.Intn(100) - 50
	if newHealth < int(p.health) && math.Abs(float64(newHealth)) > float64(p.health) {
		p.Die()
	}
	if newHealth < 0 {
		p.health -= uint8(newHealth)
	} else {
		p.health += uint8(newHealth)
	}
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
func Reproduct() interface{} {
	name := "shit"
	var sex string
	if rand.Intn(2) == 1 {
		sex = "male"
	} else {
		sex = "female"
	}
	var isRep bool
	if rand.Intn(2) == 1 {
		isRep = true
	}
	return NewParrot(name, sex, isRep, 18)
}

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
	fmt.Println("Parrot " + p.name + " is flying!")
}

func (p *Parrot) Name() string {
	return p.name
}
