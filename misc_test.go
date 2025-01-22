package jrutil

import (
	"cmp"
	"slices"
	"testing"
)

func TestSignum(t *testing.T) {
    type Data[T OrderedNumber] struct {
        x T
        expected int
    }

    data := []any {
        Data[int64]{-3, -1},
        Data[int64]{-2, -1},
        Data[int64]{-1, -1},
        Data[int64]{0, 0},
        Data[int64]{1, 1},
        Data[int64]{2, 1},
        Data[int64]{3, 1},

        Data[float64]{-3.0, -1},
        Data[float64]{-2.0, -1},
        Data[float64]{-1.0, -1},
        Data[float64]{0.0, 0},
        Data[float64]{1.0, 1},
        Data[float64]{2.0, 1},
        Data[float64]{3.0, 1},
    }

    for _, d := range data {
        switch d := d.(type) {
            case Data[int64]:
                actual := Signum(d.x)
                if actual != d.expected {
                    t.Errorf("Signum(%v): expected=%v  actual=%v", d.x, d.expected, actual)
                }
            case Data[float64]:
                actual := Signum(d.x)
                if actual != d.expected {
                    t.Errorf("Signum(%v): expected=%v  actual=%v", d.x, d.expected, actual)
                }
            default:
                t.Errorf("unexpected test case")
        }
    }
}

func TestMakeRandomBytes(t *testing.T) {
	for _, count := range []uint64{0, 1, 2, 32, 64, 128} {
		bs, err := MakeRandomBytes(count)
		if err != nil {
			t.Errorf("MakeRandomBytes: %v", err)
		}
		if uint64(len(bs)) != count {
			t.Errorf(
				"MakeRandomBytes(%v): expected_length=%v  actual_length=%v",
				count, count, len(bs))
		}
	}
}

func TestMergeSlices(t *testing.T) {
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
				"MergeSlices(%v, %v): expected=%v  actual=%v",
				d.xs, d.ys, d.expected, actual)
		}
	}
}

func TestMergeSortSlices(t *testing.T) {
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
		if !slices.Equal(actual, d.expected) {
			t.Errorf(
				"MergeSortSlices(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}
