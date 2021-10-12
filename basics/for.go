package main
import "fmt"

func superAdd2(numbers...int)int{
	total:=0
	for _,number := range numbers {
		total += number
	}
	return total
}

func main(){
	result := superAdd2(1,2,3,4,5)
	fmt.Println(result)
	
}