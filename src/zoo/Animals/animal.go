package animal

// Animal class starts the whole animal hierarchy
type Animal interface {
	LiveAnotherDay()
	Eat() uint16
	Reproduct() *Animal
	Die()
	IsAlive() bool
	Sound()
}
