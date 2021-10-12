package main
import "fmt"

func main(){	
	a:=2
	b:=&a
	a=8
	*b = 22
	fmt.Println(a,*b)

}