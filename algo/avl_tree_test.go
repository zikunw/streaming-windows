package algo

import (
	"cmp"
	"container/list"
	"reflect"
	"testing"
)

func InOrderTravAVLTree[K cmp.Ordered](n *AVLNode[K]) []K {
	results := make([]K, 0)
	current := n
	stack := list.New()
	for {
		if current != nil {
			stack.PushFront(current)
			current = current.left
		} else if stack.Len() != 0 {
			current = stack.Remove(stack.Front()).(*AVLNode[K])
			results = append(results, current.key)
			current = current.right
		} else {
			break
		}
	}
	return results
}

func TestAVLNodeInsert1(t *testing.T) {
	root := NewAVLNode(33)
	nums := []int{33, 13, 53, 9, 21, 61, 8, 11}
	for _, num := range nums {
		root = root.Insert(num)
	}

	got := InOrderTravAVLTree(root)
	expected := []int{8, 9, 11, 13, 21, 33, 53, 61}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong order of the in order traversal, expect=%v, got=%v", expected, got)
	}
}

func TestAVLNodeInsert2(t *testing.T) {
	root := NewAVLNode(8)
	nums := []int{637, 8, 19, 471, 31, 112, 882, 614, 801, 321}
	for _, num := range nums {
		root = root.Insert(num)
	}

	got := InOrderTravAVLTree(root)
	expected := []int{8, 19, 31, 112, 321, 471, 614, 637, 801, 882}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong order of the in order traversal, expect=%v, got=%v", expected, got)
	}
}

func TestAVLNodeDelete(t *testing.T) {
	root := NewAVLNode(33)
	nums := []int{33, 13, 53, 9, 21, 61, 8, 11}
	for _, num := range nums {
		root = root.Insert(num)
	}

	root = root.Delete(13)
	expected := []int{8, 9, 11, 21, 33, 53, 61}
	got := InOrderTravAVLTree(root)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong order of the in order traversal, expect=%v, got=%v", expected, got)
	}

	root = root.Delete(61)
	expected = []int{8, 9, 11, 21, 33, 53}
	got = InOrderTravAVLTree(root)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong order of the in order traversal, expect=%v, got=%v", expected, got)
	}

	root = root.Delete(9)
	expected = []int{8, 11, 21, 33, 53}
	got = InOrderTravAVLTree(root)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong order of the in order traversal, expect=%v, got=%v", expected, got)
	}
}
