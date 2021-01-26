package main

import (
	"fmt"
)

func Hello(name string, language string) string {
	
	if name == ""{
		return "Hello,world"
	}

	return languagePrefix(language) + name
}

func languagePrefix(language string) string  {
	switch language {
		case "Spanish":
			return "Hola,"
		case "French":
			return "Bonjour,"
		default:
			return "Hello,"
	}
}

func main() {

	fmt.Println(Hello("",""))
}
