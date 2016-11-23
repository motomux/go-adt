package heap

import "errors"

// ErrInvalidIndex is given when given index is invalid
var ErrInvalidIndex = errors.New("Invalid index")

// Parent returns parent of given index
func Parent(heap []int, index int) (value, i int, err error) {
	if index <= 0 || index > len(heap)-1 {
		return value, i, ErrInvalidIndex
	}

	i = (index - 1) / 2
	return heap[i], i, nil
}

// LeftChild returns left child of given index
func LeftChild(heap []int, index int) (value, i int, err error) {
	if index < 0 || index >= len(heap)/2 {
		return value, i, ErrInvalidIndex
	}

	i = index*2 + 1
	return heap[i], i, nil
}

// RightChild returns right child of given index
func RightChild(heap []int, index int) (value, i int, err error) {
	if index < 0 || index >= len(heap)/2-1 {
		return value, i, ErrInvalidIndex
	}

	i = index*2 + 2
	return heap[i], i, nil
}
