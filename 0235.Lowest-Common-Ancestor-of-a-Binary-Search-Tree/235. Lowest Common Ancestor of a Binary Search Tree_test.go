package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question235 struct {
	para235
	ans235
}

// para 是参数
// one 代表第一个参数
type para235 struct {
	one []int
	two []int
	thr []int
}

// ans 是答案
// one 代表第一个答案
type ans235 struct {
	one []int
}

func Test_Problem235(t *testing.T) {

	qs := []question235{

		{
			para235{[]int{}, []int{}, []int{}},
			ans235{[]int{}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ans235{[]int{6}},
		},

		{
			para235{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans235{[]int{2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 235------------------------\n")

	for _, q := range qs {
		_, p := q.ans235, q.para235
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		rootTwo := structures.Ints2TreeNode(p.two)
		rootThr := structures.Ints2TreeNode(p.thr)
		fmt.Printf("【output】:%v      \n", lowestCommonAncestor(rootOne, rootTwo, rootThr))
	}
	fmt.Printf("\n\n\n")
}
