package queue

//import "errors"

type LinkedQueue struct{
	front *Node
	back *Node
	size int
}

type Node struct {
	val int
	next *Node
}

func (linkedqueue *LinkedQueue) Enqueue(value int){

	newNode := Node{val: value, next: nil}

	if linkedqueue.size == 0{
		linkedqueue.front = &newNode
		linkedqueue.back = &newNode
	} else{
		linkedqueue.back.next = &newNode
		//linkedqueue.back = newNode
	}

	linkedqueue.size++

}

func (linkedqueue *LinkedQueue) Dequeue() int{

	first_value := linkedqueue.front.val

	linkedqueue.front = linkedqueue.front.next
	linkedqueue.size--

	return first_value
}

func (linkedqueue *LinkedQueue) Peek() int{

	return linkedqueue.front.val

}