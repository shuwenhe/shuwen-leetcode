package leetcode

import (
	"fmt"
	"testing"
)

type question973 struct {
	para973
	ans973
}

// para 是参数
// one 代表第一个参数
type para973 struct {
	one [][]int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans973 struct {
	one [][]int
}

func Test_Problem973(t *testing.T) {

	qs := []question973{

		{
			para973{[][]int{{1, 3}, {-2, 2}}, 1},
			ans973{[][]int{{-2, 2}}},
		},

		{
			para973{[][]int{{1, 3}, {-2, 2}}, 0},
			ans973{[][]int{{}}},
		},

		{
			para973{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
			ans973{[][]int{{3, 3}, {-2, 4}}},
		},

		{
			para973{[][]int{{0, 1}, {1, 0}}, 2},
			ans973{[][]int{{1, 0}, {0, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 973------------------------\n")

	for _, q := range qs {
		_, p := q.ans973, q.para973
		fmt.Printf("【input】:%v       【output】:%v\n", p, KClosest(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
