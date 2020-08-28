package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question98 struct {
	para98
	ans98
}

// para 是参数
// one 代表第一个参数
type para98 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans98 struct {
	one bool
}

func Test_Problem98(t *testing.T) {

	qs := []question98{

		{
			para98{[]int{10, 5, 15, structures.NULL, structures.NULL, 6, 20}},
			ans98{false},
		},

		{
			para98{[]int{}},
			ans98{true},
		},

		{
			para98{[]int{2, 1, 3}},
			ans98{true},
		},

		{
			para98{[]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}},
			ans98{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 98------------------------\n")

	for _, q := range qs {
		_, p := q.ans98, q.para98
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isValidBST(rootOne))
	}
	fmt.Printf("\n\n\n")
}
