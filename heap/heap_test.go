package heap

import "testing"

func TestParent(t *testing.T) {
	type (
		in struct {
			heap  []int
			index int
		}
		out struct {
			value, index int
			ok           bool
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{6, 0, true},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{6, 0, true},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{5, 2, true},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 4},
			out{4, 1, true},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, false},
		},
		"case6": {
			in{[]int{6, 4, 5, 1, 3, 2}, 0},
			out{0, 0, false},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, ok := Parent(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if ok != out.ok {
				t.Errorf("ok doesn't match. expected=%v, actual=%v\n", ok, out.ok)
			}
		})
	}
}

func TestLeftChild(t *testing.T) {
	type (
		in struct {
			heap  []int
			index int
		}
		out struct {
			value, index int
			ok           bool
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{1, 3, true},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{2, 5, true},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{0, 0, false},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, false},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 3},
			out{0, 0, false},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, ok := LeftChild(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if ok != out.ok {
				t.Errorf("ok doesn't match. expected=%v, actual=%v\n", ok, out.ok)
			}
		})
	}
}

func TestRightChild(t *testing.T) {
	type (
		in struct {
			heap  []int
			index int
		}
		out struct {
			value, index int
			ok           bool
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{3, 4, true},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{0, 0, false},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{0, 0, false},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, false},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 3},
			out{0, 0, false},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, ok := RightChild(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if ok != out.ok {
				t.Errorf("ok doesn't match. expected=%v, actual=%v\n", ok, out.ok)
			}
		})
	}
}
