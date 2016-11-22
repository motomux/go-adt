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
