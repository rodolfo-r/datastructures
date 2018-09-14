package doubly_test

import (
	"log"
	"testing"

	"github.com/rodolfo-r/datastructures/list/linked/doubly"
)

var empty, zero, one, two, three,
	zeroOne, oneTwo, zeroTwo, twoOne, oneZero,
	zeroOneTwo, twoOneZero *doubly.List

func init() {
	empty, zero, one, two = doubly.New(), doubly.New(), doubly.New(), doubly.New()

	zero.Append(0)
	one.Append(1)
	two.Append(2)

	zeroOne, oneTwo, zeroTwo = zero.Copy(), one.Copy(), zero.Copy()
	twoOne, oneZero := two.Copy(), one.Copy()
	zeroOne.Append(1)
	oneTwo.Append(2)
	zeroTwo.Append(2)
	twoOne.Append(1)
	oneZero.Append(0)

	zeroOneTwo, twoOneZero = zeroOne.Copy(), twoOne.Copy()
	zeroOneTwo.Append(2)
	twoOneZero.Append(0)
}

func TestAppend(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		vals []int
		want *doubly.List
	}{
		{
			name: "Append one and two to empty",
			list: empty.Copy(),
			vals: []int{1, 2},
			want: oneTwo,
		},
		{
			name: "Append one and two to zero",
			list: zero.Copy(),
			vals: []int{1, 2},
			want: zeroOneTwo,
		},
		{
			name: "Append two to zero and one",
			list: zeroOne.Copy(),
			vals: []int{2},
			want: zeroOneTwo,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < len(tc.vals); i++ {
				tc.list.Append(tc.vals[i])
			}
			if !tc.list.Equals(tc.want) {
				t.Fatalf("have %s, want %s", tc.list, tc.want)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		vals []int
		want *doubly.List
	}{
		{
			name: "Prepend two and one to empty",
			list: empty.Copy(),
			vals: []int{2, 1},
			want: oneTwo,
		},
		{
			name: "Prepend one and zero to two",
			list: two.Copy(),
			vals: []int{1, 0},
			want: zeroOneTwo,
		},
		{
			name: "Prepend zero to one and two",
			list: oneTwo.Copy(),
			vals: []int{0},
			want: zeroOneTwo,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < len(tc.vals); i++ {
				tc.list.Prepend(tc.vals[i])
			}
			if !tc.list.Equals(tc.want) {
				t.Fatalf("have %s, want %s", tc.list, tc.want)
			}
		})
	}
}

func TestFront(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		want doubly.Node
	}{
		{
			name: "Empty",
			list: empty.Copy(),
			want: doubly.Node{},
		},
		{
			name: "Zero",
			list: zero.Copy(),
			want: doubly.Node{Value: 0},
		},
		{
			name: "Zero One",
			list: zeroOne.Copy(),
			want: doubly.Node{Value: 0},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := tc.list.Front()

			if f != f {
				t.Fatalf("have %#v, want %#v", f, tc.want)
			}
		})
	}
}

func TestBack(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		want doubly.Node
	}{
		{
			name: "Empty",
			list: empty.Copy(),
			want: doubly.Node{},
		},
		{
			name: "Zero",
			list: zero.Copy(),
			want: doubly.Node{Value: 0},
		},
		{
			name: "Zero One",
			list: zeroOne.Copy(),
			want: doubly.Node{Value: 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f := tc.list.Front()

			if f != f {
				t.Fatalf("have %#v, want %#v", f, tc.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		val  int
		ok   bool
	}{
		{
			name: "Search for one in empty",
			list: empty.Copy(),
			val:  1,
			ok:   false,
		},
		{
			name: "Search for one in one",
			list: one.Copy(),
			val:  1,
			ok:   true,
		},
		{
			name: "Search for one in zero one two",
			list: zeroOneTwo.Copy(),
			val:  1,
			ok:   true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if n, ok := tc.list.Search(tc.val); ok != tc.ok {
				t.Fatalf("could not find %v in list", tc.val)
			} else if !tc.ok && !ok {
				return
			} else if n.Value != tc.val {
				log.Printf("tc.ok = %+v\n", tc.ok)
				log.Printf("ok: %v", ok)
				t.Fatalf("have %v, want %v", n.Value, tc.val)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		val  int
		want *doubly.List
	}{
		{
			name: "Delete zero from zero",
			list: zero.Copy(),
			val:  0,
			want: empty,
		},
		{
			name: "Delete zero from zero one",
			list: zeroOne.Copy(),
			val:  0,
			want: one,
		},
		{
			name: "Delete one from zero one",
			list: zeroOne.Copy(),
			val:  1,
			want: zero,
		},
		{
			name: "Delete one from zero one two",
			list: zeroOneTwo.Copy(),
			val:  1,
			want: zeroTwo,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if n, ok := tc.list.Search(tc.val); !ok {
				t.Fatalf("failed searching for %v", tc.val)
			} else {
				tc.list.Delete(n)
				if !tc.list.Equals(tc.want) {
					t.Fatalf("have %s, want %s", tc.list, tc.want)
				}
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tt := []struct {
		name string
		list *doubly.List
		want *doubly.List
	}{
		{
			name: "Empty",
			list: empty.Copy(),
			want: empty,
		},
		{
			name: "Zero",
			list: zero.Copy(),
			want: zero,
		},
		{
			name: "Zero One Two",
			list: zeroOneTwo.Copy(),
			want: twoOneZero,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Reverse()
			if !tc.list.Equals(tc.want) {
				t.Fatalf("have %s, want %s", tc.list, tc.want)
			}
		})
	}
}

func TestEach(t *testing.T) {
	zeroZeroZero := doubly.New()
	zeroZeroZero.Append(0)
	zeroZeroZero.Append(0)
	zeroZeroZero.Append(0)

	zeroTenTwenty := doubly.New()
	zeroTenTwenty.Append(0)
	zeroTenTwenty.Append(10)
	zeroTenTwenty.Append(20)

	tt := []struct {
		name string
		list *doubly.List
		fun  func(*doubly.Node)
		want *doubly.List
	}{
		{
			name: "Set each node val to 0",
			list: zeroOneTwo.Copy(),
			fun:  func(n *doubly.Node) { n.Value = 0 },
			want: zeroZeroZero,
		},
		{
			name: "Multiply each node val by 10",
			list: zeroOneTwo.Copy(),
			fun:  func(n *doubly.Node) { n.Value = n.Value.(int) * 10 },
			want: zeroTenTwenty,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Each(tc.fun)
			if !tc.list.Equals(tc.want) {
				t.Fatalf("have %s, want %s", tc.list, tc.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	zeroZeroZero := doubly.New()
	zeroZeroZero.Append(0)
	zeroZeroZero.Append(0)
	zeroZeroZero.Append(0)

	zeroTenTwenty := doubly.New()
	zeroTenTwenty.Append(0)
	zeroTenTwenty.Append(10)
	zeroTenTwenty.Append(20)

	tt := []struct {
		name         string
		list         *doubly.List
		fun          func(interface{}) interface{}
		want, mapped *doubly.List
	}{
		{
			name:   "Map each node val to 0",
			list:   zeroOneTwo.Copy(),
			fun:    func(val interface{}) interface{} { return 0 },
			want:   zeroOneTwo.Copy(),
			mapped: zeroZeroZero.Copy(),
		},
		{
			name:   "Map Multiply each node val by 10",
			list:   zeroOneTwo.Copy(),
			fun:    func(val interface{}) interface{} { return val.(int) * 10 },
			want:   zeroOneTwo,
			mapped: zeroTenTwenty.Copy(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mapped := tc.list.Map(tc.fun)

			if !tc.list.Equals(tc.want) {
				t.Fatalf("have %s, want %s", tc.list, tc.want)
			}

			if !mapped.Equals(tc.mapped) {
				t.Fatalf("have mapped %s, want %s", mapped, tc.mapped)
			}
		})
	}
}
