package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question203 struct {
	para203
	ans203
}

// para 是参数
// one 代表第一个参数
type para203 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans203 struct {
	one []int
}

func Test_Problem203(t *testing.T) {

	qs := []question203{

		{
			para203{[]int{1, 2, 3, 4, 5}, 1},
			ans203{[]int{2, 3, 4, 5}},
		},

		{
			para203{[]int{1, 2, 3, 4, 5}, 2},
			ans203{[]int{1, 3, 4, 5}},
		},

		{
			para203{[]int{1, 1, 1, 1, 1}, 1},
			ans203{[]int{}},
		},

		{
			para203{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 2},
			ans203{[]int{1, 3, 3, 3}},
		},

		{
			para203{[]int{1, 2, 3, 4, 5}, 5},
			ans203{[]int{1, 2, 3, 4}},
		},

		{
			para203{[]int{}, 5},
			ans203{[]int{}},
		},

		{
			para203{[]int{1, 2, 3, 4, 5}, 10},
			ans203{[]int{1, 2, 3, 4, 5}},
		},

		{
			para203{[]int{1}, 1},
			ans203{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 203------------------------\n")

	for _, q := range qs {
		_, p := q.ans203, q.para203
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(removeElements(structures.Ints2List(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}
