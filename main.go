package main

import "fmt"

func main() {
	name := ""
	println("Введите имя")
	fmt.Scanln(&name)
	greet(name)
}

func greet(a string) {
	fmt.Printf("Привет %s\n", a)
}
