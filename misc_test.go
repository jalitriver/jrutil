package jrutil

import (
	"cmp"
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	type Data struct {
		xs       []int
		ys       []int
		expected []int
	}

	data := []Data{
		{
			xs:       []int{},
			ys:       []int{},
			expected: []int{},
		},
		{
			xs:       []int{0},
			ys:       []int{},
			expected: []int{0},
		},
		{
			xs:       []int{},
			ys:       []int{0},
			expected: []int{0},
		},
		{
			xs:       []int{0, 1},
			ys:       []int{},
			expected: []int{0, 1},
		},
		{
			xs:       []int{},
			ys:       []int{0, 1},
			expected: []int{0, 1},
		},
		{
			xs:       []int{0, 1, 2},
			ys:       []int{},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{},
			ys:       []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{0},
			ys:       []int{1},
			expected: []int{0, 1},
		},
		{
			xs:       []int{1},
			ys:       []int{0},
			expected: []int{0, 1},
		},
		{
			xs:       []int{0, 2},
			ys:       []int{1},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{1},
			ys:       []int{0, 2},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{0, 2},
			ys:       []int{1, 3},
			expected: []int{0, 1, 2, 3},
		},
	}

	for _, d := range data {
		actual := MergeSlices(d.xs, d.ys, cmp.Less[int])
		if !slices.Equal(d.expected, actual) {
			t.Errorf(
				"Merge(%v, %v): expected=%v  actual=%v",
				d.xs, d.ys, d.expected, actual)
		}
	}
}

func TestMergeSort(t *testing.T) {
	type Data struct {
		xs       []int
		expected []int
	}

	data := []Data{
		{
			xs:       []int{},
			expected: []int{},
		},
		{
			xs:       []int{0},
			expected: []int{0},
		},
		{
			xs:       []int{0, 1},
			expected: []int{0, 1},
		},
		{
			xs:       []int{1, 0},
			expected: []int{0, 1},
		},
		{
			xs:       []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{2, 1, 0},
			expected: []int{0, 1, 2},
		},
		{
			xs:       []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, d := range data {
		actual := MergeSortSlices(d.xs, cmp.Less[int])
		if !slices.Equal(d.expected, actual) {
			t.Errorf(
				"MergeSort(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}
