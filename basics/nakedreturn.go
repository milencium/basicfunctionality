package main

import (
	"fmt"
	"strings"
)

	
func lenAndUpper2(name string)(lenght int, uppercase string){
	defer fmt.Println("done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main(){
	totalLength, uppercase := lenAndUpper2("nicos")
	fmt.Println(totalLength, uppercase)
}