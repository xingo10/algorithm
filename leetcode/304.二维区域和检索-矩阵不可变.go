/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	sums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}

	sums := make([][]int, len(matrix)+1, len(matrix)+1)
	// 二维前缀和，第一行和第一列为0
	// 先设置第一行为0
	sums[0] = make([]int, len(matrix[0])+1, len(matrix[0])+1)
	for i, row := range matrix {
        sums[i+1] = make([]int, len(row)+1, len(row)+1)
		// sum(i,j) = sum(i-1,j)+sum(i,j-1)-sum(i-1,j-1)+a(i,j)
		for j := range row {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{
		sums: sums,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// sum(x2,y2)-sum(x2,y1-1)-sum(x1-1,y2)+sum(x1-1,y1-1)
	return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

