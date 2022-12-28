package main

import "fmt"

func type_bool(){
	a,b := true ,false
	c := a && b
	fmt.Println("the value of c ",c)
	d := a || c
	fmt.Println("the value of d ",d)
	 

}