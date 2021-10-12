package main

import ("fmt")

func something(firstNumber int)(secondNumber int){
	secondNumber = firstNumber * firstNumber 
	return secondNumber
}

func main(){
	fmt.Println("go test")
	result:=something(5)
	fmt.Println(result)

}