package main

import (
	"fmt"
	"calculator-devops/internal/calc"
	
)

func main(){

	calc.InitLogger()
	calc.Logger.Println("Calculator Started")
	var a,b int
	var op string

	fmt.Println("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Println("Enter operator (+,-,*,/): ")
	fmt.Scanln(&op)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&b)

	switch op{
	case "+":
		result:=calc.Add(a,b)
		fmt.Println("Result: ",result)
		calc.Logger.Println("Result: ",result)

	case "-":
		result:=calc.Subtract(a,b)
		fmt.Println("Result: ",result)
		calc.Logger.Println("Result: ",result)

	case "*":
		result:=calc.Multiply(a,b)
		fmt.Println("Result: ",result)
		calc.Logger.Println("Result: ",result)

	case "/":
		result,err:=calc.Divide(a,b)

		if err!=nil {
			fmt.Println("Error: ",err)
			calc.Logger.Println("Error:",err)	
		} else {
			fmt.Println("Result: ",result)
			calc.Logger.Println("Result: ",result)

		}
		
	
	default:
		fmt.Println("Invalid Operator")
		calc.Logger.Println("Invalid operator entered")
	}
}