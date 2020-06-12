package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
	var num1, num2, result1 int
	num3 := 20
	fmt.Printf("Enter the values num1")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter the values num2")
	fmt.Scanf("%d", &num2)
	result1 = num1 + num2 + num3
	fmt.Printf("Result= %v", result1)
}
