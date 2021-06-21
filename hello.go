package lgwt

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(greetee string) string {
	if greetee == "" {
		greetee = "world"
	}
	return englishHelloPrefix + greetee
}
