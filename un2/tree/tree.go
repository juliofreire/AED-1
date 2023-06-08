package tree

import (
	"fmt"
)

type BstNode struct {
	left  *BstNode
	right *BstNode
	value int
}

func (node *BstNode) NewNode(value int) *BstNode {
	return &BstNode{value: value}
}

func (node *BstNode) Add(value int) {
	if node == nil {
		node = node.NewNode(value)
	}

	if value < node.value {
		if node.left == nil {
			node.left = node.NewNode(value)
		} else {
			node.left.Add(value)
		}
	} else {
		if node.right == nil {
			node.right = node.NewNode(value)
		} else {
			node.right.Add(value)
		}
	}

}

func (node *BstNode) Search(value int) bool {

	if value == node.value {
		return true
	} else if value < node.value && node.left != nil {
		return node.left.Search(value)
	} else if value > node.value && node.right != nil {
		return node.right.Search(value)
	} else {
		return false
	}
}

func (node *BstNode) Min() int {

	if node.left == nil {
		return node.value
	}
	return node.left.Min()
}

func (node *BstNode) Max() int {

	if node.right == nil {
		return node.value
	}
	return node.right.Max()
}

func (node *BstNode) PreOrder() {
	fmt.Println(node.value)
	if node.left != nil {
		node.left.PreOrder()
	}
	if node.right != nil {
		node.right.PreOrder()
	}
}

func (node *BstNode) InOrder() {

	if node.left != nil {
		node.left.InOrder()
	}
	fmt.Println(node.value)
	if node.right != nil {
		node.right.InOrder()
	}
}
func (node *BstNode) PosOrder() {

	if node.left != nil {
		node.left.PosOrder()
	}

	if node.right != nil {
		node.right.PosOrder()
	}
	fmt.Println(node.value)
}

func (node *BstNode) Remove(value int) *BstNode {
	//case 1: 0 son
	if value < node.value && node.left != nil {
		node.left = node.left.Remove(value)
	} else if value > node.value && node.right != nil {
		node.right = node.right.Remove(value)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.left == nil && node.right != nil {
			return node.right
		} else if node.left != nil && node.right != nil {
			max := node.left.Max()
			node.value = max
			node.left = node.left.Remove(max)
			return node
		}
	}
	return node
}
