/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	/*
		解题思路：
			根据题意，要求矩形最大面积=长*宽，假设长固定，那么就是宽度越大，面积越大。
			长也就是柱子的高度，如果柱子要保持高度，除非是该柱子后面的高度>=该柱子，
			否则只要出现比该柱子矮的，矩形的高就要降低。
			需要求出该柱子后面(右侧)比它矮的柱子，宽度的右侧到这里结束
			因此可以使用单调递减栈来完成。

			宽度的左侧：每次比较当前柱子i和栈顶柱子j高度，若高度j>=i，说明i柱子的宽度可以向左侧延伸，即加栈顶柱子宽度
	*/
	// 使用单调递减栈，根据题意数值都是非负数，因此最后加个0，可以帮助触发出栈
	// 否则如果heights都是单调递增的话，无法出栈，最后计算出来的面积会为初始0值
	heights = append(heights, 0)
	// 单调递减栈
	stack := make([]Rect, 0)
	area := 0

	for i := range heights {
		// 每个柱子下标需要累加宽度，因为面积是（下标左侧小于该元素的下标加右侧第一个小于该元素的下标）的宽度*高度
		// 如： [2,1,2] -> [1,1,1]，面积为3
		width := 0
		for len(stack) != 0 && heights[i] <= stack[len(stack)-1].Height {
			// 如果i这个位置的柱子高度比栈顶的矮，就说明i柱子的宽度可以加栈顶的宽度
			// 因为高的柱子可以包含矮柱子
			width += stack[len(stack)-1].Width
			area = max(area, stack[len(stack)-1].Height*width)

			stack = stack[:len(stack)-1]
		}
		// 每个柱子宽度为1
		stack = append(stack, Rect{heights[i], width + 1})
	}

	return area
}

type Rect struct {
	Height int
	Width  int
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

