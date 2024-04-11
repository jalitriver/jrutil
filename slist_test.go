package jrutil

import (
	"cmp"
	"slices"
	"testing"
)

func TestNewSList(t *testing.T) {
	if NewSList[int]() != nil {
		t.Error("new SList is not nil")
	}
}

func TestNewSListFromSlice(t *testing.T) {
	type Data struct {
		xs       []int
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       []int{},
			expected: nil,
		},
		{
			xs:       []int{0},
			expected: NewSList[int]().PushFront(0),
		},
		{
			xs:       []int{0, 1},
			expected: NewSList[int]().PushFront(1).PushFront(0),
		},
		{
			xs:       []int{0, 1, 2},
			expected: NewSList[int]().PushFront(2).PushFront(1).PushFront(0),
		},
	}

	for _, d := range data {
		actual := NewSListFromSlice(d.xs)
		if !actual.Equal(d.expected, func(x, y int) bool { return x == y }) {
			t.Errorf("NewSListFromSlice(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}

func TestSListToSlice(t *testing.T) {
	type Data struct {
		xs       *SList[int]
		expected []int
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			expected: []int{},
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			expected: []int{0},
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: []int{0, 1},
		},
		{
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: []int{0, 1, 2},
		},
	}

	for _, d := range data {
		actual := d.xs.ToSlice()
		if !slices.Equal(actual, d.expected) {
			t.Errorf("SList.ToSlice(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}

func TestSListEmpty(t *testing.T) {
	type Data struct {
		xs       *SList[int]
		expected bool
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			expected: true,
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			expected: false,
		},
	}

	for _, d := range data {
		actual := d.xs.Empty()
		if actual != d.expected {
			t.Errorf("Empty(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}

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

func TestSListPushFront(t *testing.T) {

	type Data struct {
		xs       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       NewSList[int]().PushFront(0),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			xs:       NewSList[int]().PushFront(1).PushFront(0),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSList[int]().PushFront(2).PushFront(1).PushFront(0),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
	}

	for _, d := range data {
		if !d.xs.Equal(d.expected, func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.PushFront: expected=%v  actual=%v",
				d.expected, d.xs)
		}
	}
}

func TestSListLength(t *testing.T) {

	var actual uint64
	var expected uint64

	// Empty SList should have zero length.
	xs := NewSList[int]()
	actual = xs.Length()
	expected = 0
	if actual != expected {
		t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
			xs, expected, actual)
	}

	// Add elements to the list.
	iMax := 10
	for i := 0; i < iMax; i++ {
		xs = xs.PushFront(i)
		actual = xs.Length()
		expected = uint64(i) + 1
		if actual != expected {
			t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
				xs, expected, actual)
		}
	}

	// Take elements from the list.  This does not change the length
	// of the original list.
	for i := 0; i < iMax; i++ {
		ys := xs.Take(uint64(i))
		// Verify that the length of xs has not changed.
		actual = xs.Length()
		expected = uint64(iMax)
		if actual != expected {
			t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
				xs, expected, actual)
		}
		// Verify the length of ys.
		actual = ys.Length()
		expected = uint64(i)
		if actual != expected {
			t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
				ys, expected, actual)
		}
	}

	// Drop elements from the list.  This does not change the length
	// of the original list.
	for i := 0; i < iMax; i++ {
		ys := xs.Drop(uint64(i))
		// Verify that the length of xs has not changed.
		actual = xs.Length()
		expected = uint64(iMax)
		if actual != expected {
			t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
				xs, expected, actual)
		}
		// Verify the length of ys.
		actual = ys.Length()
		expected = uint64(iMax) - uint64(i)
		if actual != expected {
			t.Fatalf("SList.Length(%v): expected=%v  actual=%v",
				ys, expected, actual)
		}
	}

}

func TestSListHead(t *testing.T) {
	var actualX, expectedX int
	var actualOK, expectedOK bool

	// Head() for an empty list should fail.
	xs := NewSList[int]()
	actualX, actualOK = xs.Head()
	expectedX, expectedOK = 0, false
	if (actualX != expectedX) || (actualOK != expectedOK) {
		t.Errorf("SList.Head(%v): expected=(%v, %v)  actual=(%v, %v)",
			xs, expectedX, expectedOK, actualX, actualOK)
	}

	// Add elements to the list and verify the head after each addition.
	for i := 0; i < 10; i++ {
		xs = xs.PushFront(i)
		actualX, actualOK = xs.Head()
		expectedX, expectedOK = i, true
		if (actualX != expectedX) || (actualOK != expectedOK) {
			t.Errorf("SList.Head(%v): expected=(%v, %v)  actual=(%v, %v)",
				xs, expectedX, expectedOK, actualX, actualOK)
		}
	}
}

func TestSListTail(t *testing.T) {
	type Data struct {
		xs       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{1}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{1, 2}),
		},
	}

	for _, d := range data {
		actual := d.xs.Tail()
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.Tail(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}

func TestSListReverse(t *testing.T) {
	type Data struct {
		xs       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{1, 0}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{2, 1, 0}),
		},
	}

	for _, d := range data {
		actual := d.xs.Reverse()
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.Reverse(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}

func TestSListDrop(t *testing.T) {
	type Data struct {
		n        uint64
		xs       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			n:        0,
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        1,
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        2,
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        0,
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			n:        1,
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        2,
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        0,
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			n:        1,
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{1}),
		},
		{
			n:        2,
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        3,
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        0,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			n:        1,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{1, 2}),
		},
		{
			n:        2,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{2}),
		},
		{
			n:        3,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			n:        4,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{}),
		},
	}

	for _, d := range data {
		actual := d.xs.Drop(d.n)
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.Drop(%v, %v): expected=%v  actual=%v",
				d.xs, d.n, d.expected, actual)
		}
	}
}

func TestSListDropUntil(t *testing.T) {
	type Data struct {
		breakPoint int
		xs         *SList[int]
		expected   *SList[int]
	}

	data := []Data{
		{
			breakPoint: -1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
		},
		{
			breakPoint: 0,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			breakPoint: 1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{2, 3, 4, 5}),
		},
		{
			breakPoint: 2,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{3, 4, 5}),
		},
		{
			breakPoint: 3,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{4, 5}),
		},
		{
			breakPoint: 4,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{5}),
		},
		{
			breakPoint: 5,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{}),
		},
	}

	for _, d := range data {
		actual := d.xs.DropUntil(func(x int) bool {
			return x > d.breakPoint
		})
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.DropUntil: breakPoint=%v  expected=%v  actual=%v",
				d.breakPoint, d.expected, actual)
		}
	}
}

func TestSListDropWhile(t *testing.T) {
	type Data struct {
		breakPoint int
		xs         *SList[int]
		expected   *SList[int]
	}

	data := []Data{
		{
			breakPoint: -1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
		},
		{
			breakPoint: 0,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{1, 2, 3, 4, 5}),
		},
		{
			breakPoint: 1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{2, 3, 4, 5}),
		},
		{
			breakPoint: 2,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{3, 4, 5}),
		},
		{
			breakPoint: 3,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{4, 5}),
		},
		{
			breakPoint: 4,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{5}),
		},
		{
			breakPoint: 5,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{}),
		},
	}

	for _, d := range data {
		actual := d.xs.DropWhile(func(x int) bool {
			return x <= d.breakPoint
		})
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.DropWhile: breakPoint=%v  expected=%v  actual=%v",
				d.breakPoint, d.expected, actual)
		}
	}
}

func TestSListTakeWhile(t *testing.T) {
	type Data struct {
		breakPoint int
		xs         *SList[int]
		expected   *SList[int]
	}

	data := []Data{
		{
			breakPoint: -1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{}),
		},
		{
			breakPoint: 0,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0}),
		},
		{
			breakPoint: 1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1}),
		},
		{
			breakPoint: 2,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			breakPoint: 3,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3}),
		},
		{
			breakPoint: 4,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4}),
		},
		{
			breakPoint: 5,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
		},
	}

	for _, d := range data {
		actual := d.xs.TakeWhile(func(x int) bool {
			return x <= d.breakPoint
		})
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.TakeWhile: breakPoint=%v  expected=%v  actual=%v",
				d.breakPoint, d.expected, actual)
		}
	}
}

func TestSListTakeUntil(t *testing.T) {
	type Data struct {
		breakPoint int
		xs         *SList[int]
		expected   *SList[int]
	}

	data := []Data{
		{
			breakPoint: -1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{}),
		},
		{
			breakPoint: 0,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0}),
		},
		{
			breakPoint: 1,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1}),
		},
		{
			breakPoint: 2,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			breakPoint: 3,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3}),
		},
		{
			breakPoint: 4,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4}),
		},
		{
			breakPoint: 5,
			xs:         NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
			expected:   NewSListFromSlice([]int{0, 1, 2, 3, 4, 5}),
		},
	}

	for _, d := range data {
		actual := d.xs.TakeUntil(func(x int) bool {
			return x > d.breakPoint
		})
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf("SList.TakeUntil: breakPoint=%v  expected=%v  actual=%v",
				d.breakPoint, d.expected, actual)
		}
	}
}

func TestSListContains(t *testing.T) {
	type Data struct {
		element  int
		xs       *SList[int]
		expected bool
	}

	data := []Data{
		{
			element:  1,
			xs:       NewSListFromSlice([]int{}),
			expected: false,
		},
		{
			element:  1,
			xs:       NewSListFromSlice([]int{0}),
			expected: false,
		},
		{
			element:  1,
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: true,
		},
		{
			element:  1,
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: true,
		},
	}

	for _, d := range data {
		actual := d.xs.Contains(func(x int) bool { return x == d.element })
		if actual != d.expected {
			t.Errorf("SList.Contains(%v, %v): expected=%v  actual=%v",
				d.xs, d.element, d.expected, actual)
		}
	}
}

func TestSListNth(t *testing.T) {
	type Data struct {
		n          uint64
		xs         *SList[int]
		expectedN  int
		expectedOK bool
	}

	data := []Data{
		{
			n:          0,
			xs:         NewSListFromSlice([]int{}),
			expectedN:  0,
			expectedOK: false,
		},
		{
			n:          0,
			xs:         NewSListFromSlice([]int{0}),
			expectedN:  0,
			expectedOK: true,
		},
		{
			n:          1,
			xs:         NewSListFromSlice([]int{0}),
			expectedN:  0,
			expectedOK: false,
		},
		{
			n:          0,
			xs:         NewSListFromSlice([]int{0, 1}),
			expectedN:  0,
			expectedOK: true,
		},
		{
			n:          1,
			xs:         NewSListFromSlice([]int{0, 1}),
			expectedN:  1,
			expectedOK: true,
		},
		{
			n:          2,
			xs:         NewSListFromSlice([]int{0, 1}),
			expectedN:  0,
			expectedOK: false,
		},
		{
			n:          0,
			xs:         NewSListFromSlice([]int{0, 1, 2}),
			expectedN:  0,
			expectedOK: true,
		},
		{
			n:          1,
			xs:         NewSListFromSlice([]int{0, 1, 2}),
			expectedN:  1,
			expectedOK: true,
		},
		{
			n:          2,
			xs:         NewSListFromSlice([]int{0, 1, 2}),
			expectedN:  2,
			expectedOK: true,
		},
		{
			n:          3,
			xs:         NewSListFromSlice([]int{0, 1, 2}),
			expectedN:  0,
			expectedOK: false,
		},
	}

	for _, d := range data {
		actualN, actualOK := d.xs.Nth(d.n)
		if (actualN != d.expectedN) || (actualOK != d.expectedOK) {
			t.Errorf("SList.Nth(%v, %v): expected=(%v, %v)  actual=(%v, %v)",
				d.xs, d.n, d.expectedN, d.expectedOK, actualN, actualOK)
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

func TestSListMergeSort(t *testing.T) {
	type Data struct {
		xs       *SList[int]
		expected *SList[int]
	}

	data := []Data{
		{
			xs:       NewSListFromSlice([]int{}),
			expected: NewSListFromSlice([]int{}),
		},
		{
			xs:       NewSListFromSlice([]int{0}),
			expected: NewSListFromSlice([]int{0}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{1, 0}),
			expected: NewSListFromSlice([]int{0, 1}),
		},
		{
			xs:       NewSListFromSlice([]int{0, 1, 2}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{2, 1, 0}),
			expected: NewSListFromSlice([]int{0, 1, 2}),
		},
		{
			xs:       NewSListFromSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}),
			expected: NewSListFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
	}

	for _, d := range data {
		actual := d.xs.MergeSort(cmp.Less[int])
		if !actual.Equal(
			d.expected,
			func(x1, x2 int) bool { return x1 == x2 }) {
			t.Errorf(
				"SList.MergeSort(%v): expected=%v  actual=%v",
				d.xs, d.expected, actual)
		}
	}
}
