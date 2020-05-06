package animals

// Animal class starts the whole animal hierarchy
type Animal interface {
	LiveAnotherDay()
	Eat(uint8)
	//Reproduct() interface{}
	Die()
	IsAlive() bool
	Sound()
	Display()
	Name() string
}
