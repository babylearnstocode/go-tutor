package main

import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("the value of p points to is: %v", *p)
	fmt.Printf("\nthe value of i is: %v", i)
	*p = 10
	fmt.Printf("\nthe value of p points to is: %v\n", *p)

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}