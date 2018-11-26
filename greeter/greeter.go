package greeter

//Greeter is a simple example interface that inherits two string return functions to be implemented
type Greeter interface {
	//Hello returns the word hello in some language.
	Hello() string
	//World returns the word world in some language.
	World() string
}
