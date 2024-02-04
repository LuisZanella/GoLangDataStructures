package main

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, got, want *linkedList) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v Want: %v", got, want)
		for got != nil && got.node != nil {
			t.Errorf("Got: %d", got.node.value)
			got = got.next
		}
		for want != nil && want.node != nil {
			t.Errorf("Want: %d", want.node.value)
			want = want.next
		}
	}
}

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		name string
		root *linkedList
		want *linkedList
	}{
		{"empty", nil, nil},
		{"nil list", &linkedList{nil, nil}, &linkedList{nil, nil}},
		{"duplicates", &linkedList{&node{0},
			&linkedList{&node{0},
				&linkedList{&node{0},
					&linkedList{&node{0},
						&linkedList{&node{0},
							&linkedList{&node{0},
								nil}}},
				}}}, &linkedList{&node{0}, nil}},
		{"sorted duplicates", &linkedList{&node{1},
			&linkedList{&node{1},
				&linkedList{&node{2},
					&linkedList{&node{10},
						&linkedList{&node{10},
							&linkedList{&node{15},
								nil}}},
				}}}, &linkedList{&node{1},
			&linkedList{&node{2},
				&linkedList{&node{10},
					&linkedList{&node{15},
						nil}}},
		}},
		{"unsorted", &linkedList{&node{15},
			&linkedList{&node{3},
				&linkedList{&node{3},
					&linkedList{&node{15},
						&linkedList{&node{5},
							&linkedList{&node{15},
								nil}}},
				}}}, &linkedList{&node{15},
			&linkedList{&node{3},
				&linkedList{&node{5},
					nil}}},
		},
		{"uniques", &linkedList{&node{1},
			&linkedList{&node{2},
				&linkedList{&node{3},
					&linkedList{&node{10},
						&linkedList{&node{5},
							&linkedList{&node{15},
								nil}}},
				}}}, &linkedList{&node{1},
			&linkedList{&node{2},
				&linkedList{&node{3},
					&linkedList{&node{10},
						&linkedList{&node{5},
							&linkedList{&node{15},
								nil}}},
				}}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := removeDups(c.root)
			assert(t, got, c.want)
		})
	}

}
