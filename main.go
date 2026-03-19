package main

func add(a, b int) int { return a + b }
func divide(a, b int) int { return a / b }
func multiply(a, b int) int { return a * b }

func main() {
	println("=== Калькулятор ===")
	println("10 + 5 =", add(10, 5))
	println("10 / 2 =", divide(10, 2))
	println("3 * 4 =", multiply(3, 4))
}
