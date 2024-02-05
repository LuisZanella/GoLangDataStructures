package main

import (
	"reflect"
	"testing"
)

func TestReturnKthFromLast(t *testing.T) {
	cases := []struct {
		name   string
		list   *LinkedList
		target int
		want   *LinkedList
	}{
		{"empty", &LinkedList{}, 0, nil},
		{"5th from last element",
			&LinkedList{1, &LinkedList{2,
				&LinkedList{3, &LinkedList{4,
					&LinkedList{5, &LinkedList{6,
						&LinkedList{7, nil}}}}}}},
			5,
			&LinkedList{3, &LinkedList{4,
				&LinkedList{5, &LinkedList{6,
					&LinkedList{7, nil}}}}}},
		{"0 target one",
			&LinkedList{1,
				&LinkedList{2,
					&LinkedList{3, nil}}},
			0,
			nil,
		},
		{"First One",
			&LinkedList{1,
				&LinkedList{2,
					&LinkedList{3, nil}}},
			3,
			&LinkedList{1,
				&LinkedList{2,
					&LinkedList{3, nil}}},
		},
		{"Last one",
			&LinkedList{1,
				&LinkedList{2,
					&LinkedList{3, nil}}},
			1,
			&LinkedList{3, nil},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := ReturnKthFromLast(c.list, c.target)
			assert(t, got, c.want)
		})
	}

}

func assert(t *testing.T, got, want *LinkedList) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %+v Want: %+v", got, want)
	}
}
