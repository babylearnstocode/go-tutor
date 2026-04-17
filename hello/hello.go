package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Tommy")
	fmt.Println(message)
}