package main

import (
	"fmt"

	"github.com/mirzaahmedov/data-structures/list"
)

func main(){
	linkedList := list.NewLinkedList()

	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(10)

	linkedList.ForEach(func(val int) {
		fmt.Println(val)
	})

	fmt.Println(linkedList.Size())
}