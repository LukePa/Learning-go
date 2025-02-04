package greetings

import "fmt"

func Hello(name string) string {
	//Return a greeting with embedded name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
