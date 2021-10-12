package main
import "fmt"


func canIDrink2(age int)bool{
	
   switch koreanAge:= age + 2;koreanAge {
   case 10:
	return false
   case 22:
	return true
	}
	return false 
	
}

func ageNumber(number int)bool{
	switch {
	case number < 18:
		return false
	case number >18:
		return true
	}
	return false
}

func main(){
	fmt.Println(canIDrink2(20))
	fmt.Println(ageNumber(25))
}