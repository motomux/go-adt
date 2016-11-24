package heap

// MaxHeapify heapifys max heap
func MaxHeapify(heap []int, index int) {
	val := heap[index]
	lv, li, lok := LeftChild(heap, index)
	rv, ri, rok := RightChild(heap, index)

	if !lok {
		return
	} else if !rok {
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
