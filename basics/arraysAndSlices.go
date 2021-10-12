package main
import "fmt"

func main(){
	//arrays, ne moran ga popunit do kraja, ostat ce mjesta
	array := [5]string{"one","one","one","one"}
	array[4] = "green"
	fmt.Println(array)
	array2 := [5]string{}
	fmt.Println(array2)
	//slices, arrays without lenght
	slice1 := []string{"element", "element"}
	slice1 = append(slice1,"element","element")
	fmt.Println(slice1)
}