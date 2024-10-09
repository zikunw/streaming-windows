package algo

import (
	"cmp"
)

// Self balance tree
type AVLNode[K cmp.Ordered] struct {
	key    K
	left   *AVLNode[K]
	right  *AVLNode[K]
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewNode[K cmp.Ordered](key K) *AVLNode[K] {
	node := &AVLNode[K]{
		key:    key,
		left:   nil,
		right:  nil,
		height: 1,
	}
	return node
}

func (n *AVLNode[K]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (y *AVLNode[K]) rightRotate() *AVLNode[K] {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1
	return x
}

func (x *AVLNode[K]) leftRotate() *AVLNode[K] {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	return y
}

func (n *AVLNode[K]) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

func (n *AVLNode[K]) Insert(key K) *AVLNode[K] {
	if n == nil {
		return NewNode(key)
	}
	if key < n.key {
		n.left = n.left.Insert(key)
	} else if key > n.key {
		n.right = n.right.Insert(key)
	} else {
		return n
	}

	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
	balance := n.getBalanceFactor()

	if balance > 1 {
		if key < n.left.key {
			return n.rightRotate()
		} else if key > n.left.key {
			n.left = n.left.leftRotate()
			return n.rightRotate()
		}
	}

	if balance < -1 {
		if n.right.getBalanceFactor() <= 0 {
			return n.leftRotate()
		} else {
			n.right = n.right.rightRotate()
			return n.leftRotate()
		}
	}

	return n
}
