package main

import ("fmt"
"unsafe")

func unsafe_sizeoff(){
	a := 100
	b := 150
	fmt.Printf("the type of a %T",a)
	fmt.Println("the size of a %d",unsafe.Sizeof(a))
	fmt.Println("the type of b ",unsafe.Sizeof(b))

}