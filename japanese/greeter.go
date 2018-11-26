package japanese

import "github.com/Bios-Marcel/golang-api-facade-example/greeter"

//GreeterJapanese is an implementation of the Greeter interface.
type GreeterJapanese struct {
	greeter.Greeter
}

//Hello returns the japanese kanji for the worl hello.
func (g *GreeterJapanese) Hello() string {
	return "こんにちは"
}

//World returns the japanese kanji for the word world.
func (g *GreeterJapanese) World() string {
	return "世界"
}
