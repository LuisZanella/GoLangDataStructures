package main

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	cases := []struct {
		name   string
		values *LinkedList
		want   *LinkedList
	}{
		{"Remove middle",
			&LinkedList{1,
				&LinkedList{2,
					&LinkedList{3,
						&LinkedList{4, nil}}}},
			&LinkedList{1,
				&LinkedList{3,
					&LinkedList{4, nil}}},
		},
	}
	for _, c := range cases {
		_, _ = DeleteNode(c.values.next)
		assert(t, c.values, c.want)
	}
}
func assert(t *testing.T, got, want *LinkedList) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %+v Want: %+v", got, want)
	}
}
