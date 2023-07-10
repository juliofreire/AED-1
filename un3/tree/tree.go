package tree

import (
	"fmt"
)

type ITree interface {
	Add(Value int) *BstNode
	Search(Value int) bool
	Min() int
	Max() int
	PreOrder()
	InOrder()
	PosOrder()
	Remove(Value int) *BstNode
	Height() int
	// IsBst() bool
	//Size() int
	RotRight() *BstNode
	RotLeft() *BstNode
	Rebalance() *BstNode
	UpdateProperties()
}

type BstNode struct {
	Left   *BstNode
	Right  *BstNode
	Value  int
	height int
	bf     int
}

func (node *BstNode) NewNode(value int) *BstNode {
	return &BstNode{Value: value, height: 0, bf: 0}
}

func (node *BstNode) Add(value int) *BstNode {
	if node == nil {
		node = node.NewNode(value)
	}

	if value < node.Value {
		if node.Left == nil {
			node.Left = node.NewNode(value)
		} else {
			node.Left = node.Left.Add(value)
		}
	} else {
		if node.Right == nil {
			node.Right = node.NewNode(value)
		} else {
			node.Right = node.Right.Add(value)
		}
	}

	node.UpdateProperties()
	return node
}

func (node *BstNode) Search(value int) bool {

	if value == node.Value {
		return true
	} else if value < node.Value && node.Left != nil {
		return node.Left.Search(value)
	} else if value > node.Value && node.Right != nil {
		return node.Right.Search(value)
	} else {
		return false
	}
}

func (node *BstNode) Min() int {

	if node.Left == nil {
		return node.Value
	}
	return node.Left.Min()
}

func (node *BstNode) Max() int {

	if node.Right == nil {
		return node.Value
	}
	return node.Right.Max()
}

func (node *BstNode) PreOrder() {
	fmt.Println(node.Value)
	if node.Left != nil {
		node.Left.PreOrder()
	}
	if node.Right != nil {
		node.Right.PreOrder()
	}
}

func (node *BstNode) InOrder() {

	if node.Left != nil {
		node.Left.InOrder()
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		node.Right.InOrder()
	}
}
func (node *BstNode) PosOrder() {

	if node.Left != nil {
		node.Left.PosOrder()
	}

	if node.Right != nil {
		node.Right.PosOrder()
	}
	fmt.Println(node.Value)
}

func (node *BstNode) Remove(value int) *BstNode {
	//TODO: check if the value exists
	if value < node.Value && node.Left != nil {
		node.Left = node.Left.Remove(value)
	} else if value > node.Value && node.Right != nil {
		node.Right = node.Right.Remove(value)
	} else {
		//case 1: 0 son
		if node.Left == nil && node.Right == nil {
			return nil
			//case 2: 1 son
		} else if node.Left != nil && node.Right == nil {
			return node.Left
		} else if node.Left == nil && node.Right != nil {
			return node.Right
			//case 3: 2 son
		} else if node.Left != nil && node.Right != nil {
			max := node.Left.Max()
			node.Value = max
			node.Left = node.Left.Remove(max)
			return node
		}
	}

	node.UpdateProperties()
	return node
}

func (node *BstNode) Height() int {
	heightL := 0
	heightR := 0

	if node.Left != nil {
		heightL = 1 + node.Left.Height()
	}
	if node.Right != nil {
		heightR = 1 + node.Right.Height()
	}

	if heightL > heightR {
		return heightL
	} else {
		return heightR
	}
}

func (node *BstNode) UpdateProperties() {
	heightL := 0
	heightR := 0

	if node.Left != nil {
		heightL = 1 + node.Left.height
	} else if node.Right != nil {
		heightR = 1 + node.Right.height
	}
	if heightL > heightR {
		node.height = 1 + heightL
	} else {
		node.height = 1 + heightR
	}

	node.bf = heightR - heightL
}

func (node *BstNode) RotRight() *BstNode {
	auxnode := node.Left
	node.Left = auxnode.Right
	auxnode.Right = node
	node.UpdateProperties()
	auxnode.UpdateProperties()
	return auxnode

}

func (node *BstNode) RotLeft() *BstNode {
	auxnode := node.Right
	node.Right = auxnode.Left
	auxnode.Left = node
	node.UpdateProperties()
	auxnode.UpdateProperties()
	return auxnode

}

// case 1: LeftLeft, bf: -2, bf: -1

func (node *BstNode) RebalanceLL() *BstNode {
	return node.RotRight()
}

// case 2 : LeftRight, bf: -2, bf: 0

func (node *BstNode) RebalanceLR() *BstNode {
	node.Left = node.Left.RotLeft()
	return node.RebalanceLL()
}

// case 3 : RightRight, bf: 2, bf: 1

func (node *BstNode) RebalanceRR() *BstNode {
	return node.RotLeft()
}

// case 4 : LeftRight, bf: 2, bf: 0

func (node *BstNode) RebalanceRL() *BstNode {
	node.Right = node.Right.RotRight()
	return node.RebalanceRR()
}

// merge all cases to rebalance in general way

func (node *BstNode) Rebalance() *BstNode {
	if node.bf <= -2 {
		if node.Left.bf == -1 {
			return node.RebalanceLL()
		} else if node.Left.bf == 0 {
			return node.RebalanceLL()
		} else {
			return node.RebalanceLR()
		}

	} else if node.bf >= 2 {
		if node.Right.bf == -1 {
			return node.RebalanceRR()
		} else if node.Right.bf == 0 {
			return node.RebalanceRR()
		} else {
			return node.RebalanceRL()
		}
	}

	return node
}
