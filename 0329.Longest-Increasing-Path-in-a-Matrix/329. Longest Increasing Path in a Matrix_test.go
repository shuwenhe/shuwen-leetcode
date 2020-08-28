package leetcode

import (
	"fmt"
	"testing"
)

type question329 struct {
	para329
	ans329
}

// para 是参数
// one 代表第一个参数
type para329 struct {
	matrix [][]int
}

// ans 是答案
// one 代表第一个答案
type ans329 struct {
	one int
}

func Test_Problem329(t *testing.T) {

	qs := []question329{

		{
			para329{[][]int{{1}}},
			ans329{1},
		},

		{
			para329{[][]int{{}}},
			ans329{0},
		},

		{
			para329{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}},
			ans329{4},
		},

		{
			para329{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}},
			ans329{4},
		},

		{
			para329{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}},
			ans329{5},
		},

		{
			para329{[][]int{{1, 5, 7}, {11, 12, 13}, {12, 13, 15}}},
			ans329{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 329------------------------\n")

	for _, q := range qs {
		_, p := q.ans329, q.para329
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestIncreasingPath(p.matrix))
	}
	fmt.Printf("\n\n\n")
}
