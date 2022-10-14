package main

import "fmt"

// Constants should improve performance of your application
// as it saves you creating the "Hello, " string instance every time Hello is called.
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Salut, "

// Capital letter function is the public function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

// the prefix will create a variable in the function
// and it's assigned a "zero" value
// it depends on the type tho, so if the type is int, it will be assigned 0
func greetingPrefix(language string) (prefix string) {
	// and this ^^ is a private function
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	// You can return whatever it's set to by just calling `return`
	// rather than `return prefix`
	return

}

func main() {
	fmt.Println(Hello("Peggy", ""))
}
