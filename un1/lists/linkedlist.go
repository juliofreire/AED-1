package lists

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func (linkedlist *LinkedList) Add(value int)  {

	newNode := Node{val: value, next: nil}

	if linkedlist.head == nil {
		linkedlist.head = &newNode

	} else {
		aux := linkedlist.head
		cur := aux
		for aux != nil {
			cur = aux
			aux = aux.next
		}
		cur.next = &newNode
	}
	linkedlist.size++

}


func (linkedlist *LinkedList) AddOnIndex(value int, index int) error {
	
	newNode := Node{val: value, next: nil}

	if index >= 0 && index <= linkedlist.size{
		if index == 0{
			if linkedlist.head == nil{
				linkedlist.head = &newNode
			} else{
			newNode.next = linkedlist.head
			linkedlist.head = &newNode
		}
		} else {
		aux := linkedlist.head
		cur := aux
		for i := 0; i <= index-1; i++{
			cur = aux
			aux = aux.next
		}
		newNode.next = aux
		cur.next = &newNode
		} 
		linkedlist.size++
		return nil
	}else if index < 0{
		return errors.New("It's impossible add a value on a negative index")

	}else{
		return errors.New("It's impossible add a value on a index that exceeds the size of Linked List")
	}
}

func (linkedlist *LinkedList) RemoveOnIndex(index int) error {
	
	if index >= 0 && index <= linkedlist.size{
		if index == 0{
			linkedlist.head = linkedlist.head.next
		}else {
		aux := linkedlist.head
		cur := aux
		for i := 0; i < index; i++{
			cur = aux
			aux = aux.next
		}
		cur.next = aux.next
		}
		linkedlist.size--
		return nil
		
	} else if index < 0{
		return errors.New("It's impossible remove a value on a negative index")

	}else{
		return errors.New("It's impossible remove a value on a index that exceeds the size of Linked List")
	}
}

func (linkedlist *LinkedList) PrintList() error {

	if linkedlist.size >= 1 {
		aux := linkedlist.head
		cur := aux
		idx := 0
		for aux != nil {
			cur = aux
			fmt.Printf("index: %v, valor %v\n", idx ,cur.val)
			aux = aux.next
			idx++
		}
		return nil
	} else {
		return errors.New("There isn't data on the Linked List")
	}
}

func (linkedlist *LinkedList) Size() int {
	return linkedlist.size
}
