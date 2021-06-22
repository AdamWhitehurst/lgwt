package lgwt

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(n, l string) string {
	if n == "" {
		n = "world"
	}

	return greetingPrefix(l) + n
}

func greetingPrefix(l string) (p string) {
	switch l {
	case spanish:
		p = spanishHelloPrefix
	case french:
		p = frenchHelloPrefix
	default:
		p = englishHelloPrefix
	}

	return
}
