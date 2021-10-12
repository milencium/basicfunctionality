package main
import (
	"fmt"
	"something"
	"strings"
)


func multiply(a int,b int) int{
	return a*b
}
func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func arrayOfWords(words ...string){
	fmt.Println(words)
}

var outside bool = true
func main(){
	const name string =  "nico"
	var name2 string =  "nico"
	name2 = "lynn"
	//faster way, go will guess the type
	nameMy:="Mile"
	//this will work since we are using exported function
	something.SayHello()
	//this won't work since we are using private function
	//something.sayBye() 	
	fmt.Println(name2)
	fmt.Println(nameMy)

	fmt.Println(multiply(2,2))
	//direct use of the function
	fmt.Println(lenAndUpper("mynameis"))
	//use of the function where we first save it to variable
	totalLenght, upperName := lenAndUpper("mynameis")
	fmt.Println(totalLenght, upperName)
	arrayOfWords("nico", "lynn", "mile", "yeah")
}