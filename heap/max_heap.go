package heap

// MaxHeapify heapifys max heap
func MaxHeapify(heap []int, index int) {
	val := heap[index]
	lv, li, lerr := LeftChild(heap, index)
	rv, ri, rerr := RightChild(heap, index)

	if lerr == ErrInvalidIndex {
		return
	} else if rerr == ErrInvalidIndex {
		if val < lv {
			heap[index], heap[li] = heap[li], heap[index]
			MaxHeapify(heap, li)
		}
	} else {
		if val > lv && val > rv {
			return
		} else if lv > rv {
			heap[index], heap[li] = heap[li], heap[index]
			MaxHeapify(heap, li)
		} else {
			heap[index], heap[ri] = heap[ri], heap[index]
			MaxHeapify(heap, ri)
		}
	}
	return
}
