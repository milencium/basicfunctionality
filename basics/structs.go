package main
import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main(){
	//structs
	favFood := []string{"kimchi","remen"}
	human := person{"name",22,favFood}
	human2 := person{name:"milencio",age: 22 , favFood:favFood}
	fmt.Println(human.name)
	
	fmt.Println(human2.name)


}