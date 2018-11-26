package main

import (
	"fmt"

	"github.com/Bios-Marcel/golang-api-facade-example/english"
	"github.com/Bios-Marcel/golang-api-facade-example/greeter"
	"github.com/Bios-Marcel/golang-api-facade-example/japanese"
)

func main() {
	var englishGreeter greeter.Greeter = &english.GreeterEnglish{}
	printGreeting(englishGreeter)

	var japaneseGreeter greeter.Greeter = &japanese.GreeterJapanese{}
	printGreeting(japaneseGreeter)
}

func printGreeting(greeter greeter.Greeter) {
	fmt.Println(greeter.Hello(), greeter.World())
}
