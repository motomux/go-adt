package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeapify(t *testing.T) {
	type (
		in struct {
			heap  []int
			index int
		}
	)

	tests := map[string]struct {
		in
		post []int
	}{
		"case1": {
			in{[]int{6, 2, 4, 3, 1, 0}, 1},
			[]int{6, 3, 4, 2, 1, 0},
		},
		"case2": {
			in{[]int{6, 5, 4, 3, 1, 0}, 1},
			[]int{6, 5, 4, 3, 1, 0},
		},
		"case3": {
			in{[]int{1, 6, 5, 3, 4, 2}, 0},
			[]int{6, 4, 5, 3, 1, 2},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, post := test.in, test.post
			MaxHeapify(in.heap, in.index)
			if !reflect.DeepEqual(post, in.heap) {
				t.Errorf("expected: %v, actual: %v", post, in.heap)
			}
		})
	}
}
