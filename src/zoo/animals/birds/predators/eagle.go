package bpd

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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
func (e *Eagle) LiveAnotherDay() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newHealth := r.Intn(100) - 50
	if newHealth < int(e.health) && math.Abs(float64(newHealth)) > float64(e.health) {
		e.Die()
	}
	if newHealth < 0 {
		e.health -= uint8(newHealth)
	} else {
		e.health += uint8(newHealth)
	}
}

// Eat eats a certain amount of food
func (e *Eagle) Eat(amount uint8) {
	e.hunger += amount
}

// Reproduce reproduces another Eagle
func (e *Eagle) Reproduce() *Eagle {
	return new(Eagle)
}

// Die symbolizes Eagles death
func (e *Eagle) Die() {
	e.isAlive = false
}

// IsAlive returns fact of Eagle's living
func (e *Eagle) IsAlive() bool {
	return e.isAlive
}

// Sound makes some sound
func (e *Eagle) Sound() {
	fmt.Println("Eagle goes brrrr!")
}

// Reproduct returns a new instance os a class
func Reproduct() interface{} {
	name := "predator shit"
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
	return NewEagle(name, sex, isRep, 18)
}

// Display Displays info about a eagle
func (e *Eagle) Display() {
	fmt.Println(*e)
}

// Fly ...
func (e *Eagle) Fly() {
	fmt.Println("Eagle " + e.name + " is flying")
}

// NewEagle ...
func NewEagle(name, sex string, isRep bool, repMinAge uint8) *Eagle {
	return &Eagle{
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

func (e *Eagle) Name() string {
	return e.name
}
