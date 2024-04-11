package jrutil

import (
	"cmp"
	"testing"
)

func TestSListEqual(t *testing.T) {

	type Data struct {
		xs       *SList[int]
		ys       *SList[int]
		expected bool
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{}),
			expected: true,
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			ys:       NewSListFromSlice([]int{}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{0}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			ys:       NewSListFromSlice([]int{0}),
			expected: true,
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{0, 1}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{0}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			ys:       NewSListFromSlice([]int{0, 1}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{0, 1}),
			expected: true,
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{1, 0}),
			expected: false,
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{0, 0}),
			expected: false,
		},
	}

	for _, d := range data {
		actual := d.xs.Equal(d.ys, func(x, y int) bool { return x == y })
		if actual != d.expected {
			t.Errorf(
				"SList.Equal(%v, %v): expected=%v  actual=%v",
				d.xs, d.ys, d.expected, actual)
		}
	}
}

func TestSListMerge(t *testing.T) {

	type Data struct {
		xs       *SList[int]
		ys       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			ys:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			ys:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			ys:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{}),
			ys:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			ys:       NewSListFromSlice([]int{1}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{1}),
			ys:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 2}),
			ys:       NewSListFromSlice([]int{1}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{1}),
			ys:       NewSListFromSlice([]int{0, 2}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 2}),
			ys:       NewSListFromSlice([]int{1, 3}),
			expected: NewSListFromSlice([]int{0, 1, 2, 3}),
		},
	}

	for _, d := range data {
		actual := d.xs.Merge(d.ys, cmp.Less)
		if !actual.Equal(d.expected, func(x, y int) bool { return x == y }) {
			t.Errorf(
				"SList.Merge(%v, %v): expected=%v  actual=%v",
				d.xs, d.ys, d.expected, actual)
		}
	}

}
