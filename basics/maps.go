package main
import "fmt"

func main(){
	//maps
	milencio := map[string]string {"name":"milencio", "age" : "15"}
	fmt.Println(milencio)
	for key,value := range milencio{
		fmt.Println(key,value)
	}
}