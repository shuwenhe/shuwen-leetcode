package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question968 struct {
	para968
	ans968
}

// para 是参数
// one 代表第一个参数
type para968 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans968 struct {
	one int
}

func Test_Problem968(t *testing.T) {

	qs := []question968{

		{
			para968{[]int{0, 0, structures.NULL, 0, 0}},
			ans968{1},
		},

		{
			para968{[]int{0, 0, structures.NULL, 0, structures.NULL, 0, structures.NULL, structures.NULL, 0}},
			ans968{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 968------------------------\n")

	for _, q := range qs {
		_, p := q.ans968, q.para968
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", minCameraCover(root))
	}
	fmt.Printf("\n\n\n")
}
