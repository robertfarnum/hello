package main

import "fmt"

// Public Language Constants
const English = "english"
const Spanish = "spanish"
const French = "french"

// Private Hello Prefix translations
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

// Private World translations
const worldEnglish = "World"
const worldSpanish = "Mundo"
const worldFrench = "Munde"

// Get the language tranlation prefix
func getPrefix(language string) string {
	switch language {
	case English:
		return helloEnglishPrefix
	case Spanish:
		return helloSpanishPrefix
	case French:
		return helloFrenchPrefix
	}

	return helloEnglishPrefix
}

// Get the "World" language tranlation
func getWorld(language string) string {
	switch language {
	case English:
		return worldEnglish
	case Spanish:
		return worldSpanish
	case French:
		return worldFrench
	}

	return worldEnglish
}

// Return the Greeting with name or "World" based on tranlation language
func Hello(name string, language string) string {
	if name == "" {
		name = getWorld(language)
	}

	return getPrefix(language) + name
}

// Main function
func main() {
	fmt.Println(Hello("Robert", English))
}
