package main

import (
	//"fmt"
	//"fmt"

	//"./lists"
	"fmt"

	"./queue"
)

func main (){

	// tests of arraylist
	//arraylist := lists.ArrayList{}
	//arraylist.Init(5)
//
	//arraylist.Add(1)
	//arraylist.Add(2)
	//arraylist.Add(3)
	//arraylist.Add(4)
	//arraylist.Add(5)
	//
//
	////a, ar := arraylist.Get(3)
	////fmt.Println(a,ar)
//
	////s := arraylist.Size()
	//
	//arraylist.RemoveOnIndex(4)
	//s, c := arraylist.Size()
	//fmt.Println("tamanho",s, "cap:", c)
	//
	//a, ar := arraylist.Get(3)
	//fmt.Println("get", a, ar)
	
	//tests of linkedlist

	llq := queue.LinkedQueue{}

	llq.Enqueue(3)
	llq.Dequeue()
	
	llq.Enqueue(7)
	llq.Dequeue()
	llq.Enqueue(6)
	llq.Enqueue(9)

	a := llq.Peek()

	fmt.Println(a)

	fmt.Println("resto",9%5)

}