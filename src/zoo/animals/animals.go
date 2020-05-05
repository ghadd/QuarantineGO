package animals

// Animal class starts the whole animal hierarchy
type Animal interface {
	LiveAnotherDay()
	Eat(uint8)
	// Reproduct() *Animal
	Die()
	IsAlive() bool
	Sound()
	Display()
}
