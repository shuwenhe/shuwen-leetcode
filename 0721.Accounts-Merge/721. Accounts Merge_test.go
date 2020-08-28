package leetcode

import (
	"fmt"
	"testing"
)

type question721 struct {
	para721
	ans721
}

// para 是参数
// one 代表第一个参数
type para721 struct {
	w [][]string
}

// ans 是答案
// one 代表第一个答案
type ans721 struct {
	one [][]string
}

func Test_Problem721(t *testing.T) {

	qs := []question721{

		{
			para721{[][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"}}},
			ans721{[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"}}},
		},

		{
			para721{[][]string{{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
				{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"}}},
			ans721{[][]string{{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"}}},
		},

		{
			para721{[][]string{{"David", "David0@m.co", "David4@m.co", "David3@m.co"},
				{"David", "David5@m.co", "David5@m.co", "David0@m.co"},
				{"David", "David1@m.co", "David4@m.co", "David0@m.co"},
				{"David", "David0@m.co", "David1@m.co", "David3@m.co"},
				{"David", "David4@m.co", "David1@m.co", "David3@m.co"}}},
			ans721{[][]string{{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"}}},
		},

		{
			para721{[][]string{}},
			ans721{[][]string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 721------------------------\n")

	for _, q := range qs {
		_, p := q.ans721, q.para721
		fmt.Printf("【input】:%v       【output】:%v\n", p, accountsMerge(p.w))
	}
	fmt.Printf("\n\n\n")
}
