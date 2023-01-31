package main

import "fmt"

const GREET_STRING_PREFIX = "Hello"

func NamelyGreet(name string) string {

	if name == "" {
		return fmt.Sprintf("%s, World!", GREET_STRING_PREFIX)
	} else {
		return fmt.Sprintf("%s, %s!", GREET_STRING_PREFIX, name)
	}
}
