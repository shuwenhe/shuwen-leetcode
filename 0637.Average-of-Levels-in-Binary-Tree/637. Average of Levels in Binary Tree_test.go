package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question637 struct {
	para637
	ans637
}

// para 是参数
// one 代表第一个参数
type para637 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans637 struct {
	one [][]int
}

func Test_Problem637(t *testing.T) {

	qs := []question637{

		{
			para637{[]int{}},
			ans637{[][]int{}},
		},

		{
			para637{[]int{1}},
			ans637{[][]int{{1}}},
		},

		{
			para637{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans637{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 637------------------------\n")

	for _, q := range qs {
		_, p := q.ans637, q.para637
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", averageOfLevels(root))
	}
	fmt.Printf("\n\n\n")
}
