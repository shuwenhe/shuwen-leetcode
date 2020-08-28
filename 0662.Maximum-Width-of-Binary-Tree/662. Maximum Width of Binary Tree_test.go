package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question662 struct {
	para662
	ans662
}

// para 是参数
// one 代表第一个参数
type para662 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans662 struct {
	one int
}

func Test_Problem662(t *testing.T) {

	qs := []question662{

		{
			para662{[]int{1, 1, 1, 1, 1, 1, 1, structures.NULL, structures.NULL, structures.NULL, 1, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 2, 2, 2, 2, 2, 2, 2, structures.NULL, 2, structures.NULL, structures.NULL, 2, structures.NULL, 2}},
			ans662{8},
		},

		{
			para662{[]int{1, 1, 1, 1, structures.NULL, structures.NULL, 1, 1, structures.NULL, structures.NULL, 1}},
			ans662{8},
		},

		{
			para662{[]int{}},
			ans662{0},
		},

		{
			para662{[]int{1}},
			ans662{1},
		},

		{
			para662{[]int{1, 3, 2, 5, 3, structures.NULL, 9}},
			ans662{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 662------------------------\n")

	for _, q := range qs {
		_, p := q.ans662, q.para662
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", widthOfBinaryTree(root))
	}
	fmt.Printf("\n\n\n")
}
