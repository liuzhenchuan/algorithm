package array

import "sort"

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum

答题思路：
遍历第一个数，转化为找剩下两数之和为第一个数的相反数，
再通过前后指针同时向中间遍历，找到和为第一个数的相反数，得到答案
去重处理，第一个数和第二个数遍历时，遇到重复继续往前走

1、先对数组进行排序
2、遍历第一个数，如果后一个数和前一个重复了，继续向后遍历
3、第三个数为最后一个数，则问题转化为找两个数之和为第一个数的相反数
4、从第二个数开始遍历，起始位置为第一个数的下一个开始，到最后
5、遇到重复到继续向下遍历
6、第二个数小与第三个数，且两数之后大与目标值，第三个数向前找
7、两数相等了就结束，如果找到值就放到结果中
8、返回结果
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
