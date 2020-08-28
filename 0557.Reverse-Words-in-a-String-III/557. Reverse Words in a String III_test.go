package leetcode

import (
	"fmt"
	"testing"
)

type question557 struct {
	para557
	ans557
}

// para 是参数
// one 代表第一个参数
type para557 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans557 struct {
	one string
}

func Test_Problem557(t *testing.T) {

	qs := []question557{

		{
			para557{"Let's take LeetCode contest"},
			ans557{"s'teL ekat edoCteeL tsetnoc"},
		},

		{
			para557{""},
			ans557{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 557------------------------\n")

	for _, q := range qs {
		_, p := q.ans557, q.para557
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseWords(p.s))
	}
	fmt.Printf("\n\n\n")
}
