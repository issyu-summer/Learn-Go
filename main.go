package main

import "fmt"

type User struct {
	name string
	age int
}

func (receiver *User) Test()  {
	fmt.Printf("%p,%v\n",receiver,receiver)
}
func main() {

}