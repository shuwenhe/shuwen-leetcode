package leetcode

import (
	"fmt"
	"testing"
)

type question131 struct {
	para131
	ans131
}

// para 是参数
// one 代表第一个参数
type para131 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans131 struct {
	one [][]string
}

func Test_Problem131(t *testing.T) {

	qs := []question131{

		{
			para131{"aab"},
			ans131{[][]string{{"aa", "b"}, {"a", "a", "b"}}},
		},

		{
			para131{"bb"},
			ans131{[][]string{{"b", "b"}, {"bb"}}},
		},

		{
			para131{"efe"},
			ans131{[][]string{{"e", "f", "e"}, {"efe"}}},
		},

		{
			para131{"abbab"},
			ans131{[][]string{{"a", "b", "b", "a", "b"}, {"a", "b", "bab"}, {"a", "bb", "a", "b"}, {"abba", "b"}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 131------------------------\n")

	for _, q := range qs {
		_, p := q.ans131, q.para131
		fmt.Printf("【input】:%v       【output】:%v\n", p, partition131(p.s))
	}
	fmt.Printf("\n\n\n")
}
