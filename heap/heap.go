package heap

// Parent returns parent of given index
func Parent(heap []int, index int) (value, i int, ok bool) {
	if index <= 0 || index > len(heap)-1 {
		return value, i, false
	}

	i = (index - 1) / 2
	return heap[i], i, true
}

// LeftChild returns left child of given index
func LeftChild(heap []int, index int) (value, i int, ok bool) {
	if index < 0 || index >= len(heap)/2 {
		return value, i, false
	}

	i = index*2 + 1
	return heap[i], i, true
}

// RightChild returns right child of given index
func RightChild(heap []int, index int) (value, i int, ok bool) {
	if index < 0 || index >= len(heap)/2-1 {
		return value, i, false
	}

	i = index*2 + 2
	return heap[i], i, true
}
