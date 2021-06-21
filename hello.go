package lgwt

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(greetee string) string {
	return fmt.Sprintf("Hello, %s", greetee)
}
