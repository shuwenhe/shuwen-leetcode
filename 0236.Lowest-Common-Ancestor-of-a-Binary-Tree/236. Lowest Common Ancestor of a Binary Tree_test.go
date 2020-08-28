package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question236 struct {
	para236
	ans236
}

// para 是参数
// one 代表第一个参数
type para236 struct {
	one []int
	two []int
	thr []int
}

// ans 是答案
// one 代表第一个答案
type ans236 struct {
	one []int
}

func Test_Problem236(t *testing.T) {

	qs := []question236{

		{
			para236{[]int{}, []int{}, []int{}},
			ans236{[]int{}},
		},

		{
			para236{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, []int{1}},
			ans236{[]int{3}},
		},

		{
			para236{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, []int{4}},
			ans236{[]int{5}},
		},

		{
			para236{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{8}},
			ans236{[]int{6}},
		},

		{
			para236{[]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}, []int{2}, []int{4}},
			ans236{[]int{2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 236------------------------\n")

	for _, q := range qs {
		_, p := q.ans236, q.para236
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		rootTwo := structures.Ints2TreeNode(p.two)
		rootThr := structures.Ints2TreeNode(p.thr)
		fmt.Printf("【output】:%v      \n", lowestCommonAncestor236(rootOne, rootTwo, rootThr))
	}
	fmt.Printf("\n\n\n")
}
