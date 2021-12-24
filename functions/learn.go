package functions

import "fmt"

func add(x int, y int) int {
	return x + y
}

func convert(title, email string) (string, string) {
	return title, email
}
func Learn() {
	fmt.Println(add(10, 5))
	fmt.Println(convert("Hello", "hello@gmail.com"))
}
