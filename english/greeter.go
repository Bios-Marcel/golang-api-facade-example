package english

import "github.com/Bios-Marcel/golang-api-facade-example/greeter"

//GreeterEnglish is an implementation of the Greeter interface.
type GreeterEnglish struct {
	greeter.Greeter
}

//Hello returns the word Hello.
func (g *GreeterEnglish) Hello() string {
	return "Hello"
}

//World returns the word world.
func (g *GreeterEnglish) World() string {
	return "World"
}
