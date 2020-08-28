package leetcode

import (
	"fmt"
	"testing"
)

type question885 struct {
	para885
	ans885
}

// para 是参数
// one 代表第一个参数
type para885 struct {
	R  int
	C  int
	r0 int
	c0 int
}

// ans 是答案
// one 代表第一个答案
type ans885 struct {
	one [][]int
}

func Test_Problem885(t *testing.T) {

	qs := []question885{

		{
			para885{1, 4, 0, 0},
			ans885{[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		},

		{
			para885{5, 6, 1, 4},
			ans885{[][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4},
				{3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1},
				{1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 885------------------------\n")

	for _, q := range qs {
		_, p := q.ans885, q.para885
		fmt.Printf("【input】:%v       【output】:%v\n", p, spiralMatrixIII(p.R, p.C, p.r0, p.c0))
	}
	fmt.Printf("\n\n\n")
}
