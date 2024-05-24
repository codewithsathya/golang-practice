package main

import (
	"fmt"
	"golang-practice/genericlist"
)

func main() {
	glist := genericlist.New[string]()
	glist.Insert("Hello")
	glist.Insert("world")
	glist.Insert("there")
	fmt.Print(glist.Get(2))
	fmt.Print(glist.Get(-1))
}