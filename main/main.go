package main

import (
	"DSA/DS/linkedlist"
	"fmt"
)

func main() {
	l := linkedlist.NewLinkedList()

	l.Insert(10)
	l.Insert(20)
	l.Insert(30)
	fmt.Println(l)

}
