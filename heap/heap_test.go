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
			err          error
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{6, 0, nil},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{6, 0, nil},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{5, 2, nil},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 4},
			out{4, 1, nil},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, ErrInvalidIndex},
		},
		"case6": {
			in{[]int{6, 4, 5, 1, 3, 2}, 0},
			out{0, 0, ErrInvalidIndex},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, err := Parent(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if err != out.err {
				t.Errorf("err doesn't match. expected=%v, actual=%v\n", err, out.err)
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
			err          error
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{1, 3, nil},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{2, 5, nil},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{0, 0, ErrInvalidIndex},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, ErrInvalidIndex},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 3},
			out{0, 0, ErrInvalidIndex},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, err := LeftChild(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if err != out.err {
				t.Errorf("err doesn't match. expected=%v, actual=%v\n", err, out.err)
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
			err          error
		}
	)

	tests := map[string]struct {
		in
		out
	}{
		"case1": {
			in{[]int{6, 4, 5, 1, 3, 2}, 1},
			out{3, 4, nil},
		},
		"case2": {
			in{[]int{6, 4, 5, 1, 3, 2}, 2},
			out{0, 0, ErrInvalidIndex},
		},
		"case3": {
			in{[]int{6, 4, 5, 1, 3, 2}, 5},
			out{0, 0, ErrInvalidIndex},
		},
		"case4": {
			in{[]int{6, 4, 5, 1, 3, 2}, 6},
			out{0, 0, ErrInvalidIndex},
		},
		"case5": {
			in{[]int{6, 4, 5, 1, 3, 2}, 3},
			out{0, 0, ErrInvalidIndex},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, out := test.in, test.out

			value, index, err := RightChild(in.heap, in.index)

			if value != out.value {
				t.Errorf("value doesn't match. expected=%d, actual=%d\n", value, out.value)
			}
			if index != out.index {
				t.Errorf("index doesn't match. expected=%d, actual=%d\n", index, out.index)
			}
			if err != out.err {
				t.Errorf("err doesn't match. expected=%v, actual=%v\n", err, out.err)
			}
		})
	}
}
