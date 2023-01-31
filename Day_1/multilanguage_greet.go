package main

import "fmt"

func MultilanguageGreet(name, langCode string) string {
	var message string
	switch langCode {
	case "EE":
		message = fmt.Sprintf("Tere tulemast, %s!", name)
	case "AZ":
		message = fmt.Sprintf("Salam, %s!", name)
	case "EN":
		message = fmt.Sprintf("Hello, %s!", name)
	}
	return message
}
