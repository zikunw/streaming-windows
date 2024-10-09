package algo

import "cmp"

// Self balance tree
type AVLNode[K cmp.Ordered] struct {
	key    K
	left   *AVLNode[K]
	right  *AVLNode[K]
	height int

	// for order statistic tree
	size int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewAVLNode[K cmp.Ordered](key K) *AVLNode[K] {
	node := &AVLNode[K]{
		key:    key,
		left:   nil,
		right:  nil,
		height: 1,
		size:   1,
	}
	return node
}

func (n *AVLNode[K]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode[K]) getSize() int {
	if n == nil {
		return 0
	}
	return n.size
}

func (y *AVLNode[K]) rightRotate() *AVLNode[K] {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	y.size = y.left.getSize() + y.right.getSize() + 1
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1
	x.size = x.left.getSize() + x.right.getSize() + 1
	return x
}

func (x *AVLNode[K]) leftRotate() *AVLNode[K] {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(x.left.getHeight(), x.right.getHeight()) + 1
	x.size = x.left.getSize() + x.right.getSize() + 1
	y.height = max(y.left.getHeight(), y.right.getHeight()) + 1
	y.size = y.left.getSize() + y.right.getSize() + 1
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
		return NewAVLNode(key)
	}
	if key < n.key {
		n.left = n.left.Insert(key)
	} else if key > n.key {
		n.right = n.right.Insert(key)
	} else {
		return n
	}

	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
	n.size = 1 + n.left.getSize() + n.right.getSize()
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

func (n *AVLNode[K]) getNodeWithMin() *AVLNode[K] {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *AVLNode[K]) Delete(key K) *AVLNode[K] {
	if n == nil {
		return n
	}
	if key < n.key {
		n.left = n.left.Delete(key)
	} else if key > n.key {
		n.right = n.right.Delete(key)
	} else {
		if n.left == nil || n.right == nil {
			temp := n.left
			if temp == nil {
				temp = n.right
			}
			if temp == nil {
				n = nil
			} else {
				*n = *temp
			}
		} else {
			temp := n.right.getNodeWithMin()
			n.key = temp.key
			n.right = n.right.Delete(temp.key)
		}
	}
	if n == nil {
		return n
	}

	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
	n.size = 1 + n.left.getSize() + n.right.getSize()
	balance := n.getBalanceFactor()

	if balance > 1 {
		if n.left.getBalanceFactor() >= 0 {
			return n.rightRotate()
		} else {
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
