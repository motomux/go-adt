package stack

import "testing"

func TestPush(t *testing.T) {
	tests := map[string]struct {
		in   []int
		list []int
	}{
		"case1": {
			in:   []int{1, 2, 3, 4, 5},
			list: []int{5, 4, 3, 2, 1},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, list := test.in, test.list
			stack := NewStack()

			for _, num := range in {
				stack.Push(num)
			}

			var i int
			for e := stack.list.Front(); e != nil; e = e.Next() {
				if e.Value != list[i] {
					t.Errorf("expected: %v, actual: %v, index: %d\n", list[i], e.Value, i)
				}
				i++
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := map[string]struct {
		in   []int
		list []int
	}{
		"case1": {
			in:   []int{1, 2, 3, 4, 5},
			list: []int{5, 4, 3, 2, 1},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			in, list := test.in, test.list
			stack := NewStack()

			for _, num := range in {
				stack.Push(num)
			}

			for i := 0; i < len(list); i++ {
				if v := stack.Pop(); v != list[i] {
					t.Errorf("expected: %v, actual: %v, index: %d\n", list[i], v, i)
				}
			}
		})
	}
}
