package main

import "fmt"

var a, b, c, d float64

func main() {
	//1
	//2
	fmt.Print("Введите значение для a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение для b: ")
	fmt.Scan(&b)
	fmt.Print("Введите значение для c: ")
	fmt.Scan(&c)
	fmt.Print("Введите значение для d: ")
	fmt.Scan(&d)

	result := calculateAverage(a, b, c, d)
	fmt.Printf("Среднее значение: %.2f\n", result)

}
func calculateAverage(a float64, b float64, c float64, d float64) (sred float64) {
	sred = (a + b + c + d) / 4
	return sred
}
