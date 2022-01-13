package main

import "fmt"

func main() {
	fmt.Println("五五分爱情法则")
	for i := 5;i <= 10; i++{
		u := 10 - i
		fmt.Printf("%d x %d = %d \n",i,u,i*u)
	}
	fmt.Println("Copyright © Jennifer Zhang")
}
