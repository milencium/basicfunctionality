package main
import "fmt"

func canIDrink(age int)bool{
	//koreanAge := age + 2; 
	if koreanAge:=age+2; koreanAge < 18{
		return false
	}
	return true
}


func main(){
	canIDrink(15)
	fmt.Println(canIDrink(22))
}