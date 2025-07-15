package main

import "fmt"

// 爬楼梯问题
func main() {
	// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

	n := 5
	resolve01(n)

	fmt.Println()
	resolve02(n)

	fmt.Println()
	resolve03(n)

}

func resolve03(n int) {
	// 空间优化DP解决爬楼梯问题
	// Space optimization DP solves the problem of climbing stairs

	var res int = dynamicClimbingOptimized(n)

	fmt.Printf("空间优化DP解决: n = %d 阶楼梯，一共有 %d 种爬楼梯方法。\n", n, res)
}

func dynamicClimbingOptimized(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}

	// 初始两个变量作为n=1和n=2的结果
	a := 1
	b := 2

	// 根据当前结果不断往后计算，将dp的空间复杂度O(n)缩小为O(1)
	for i := 2; i < n; i++ {
		//temp := a
		//a = b
		//b = temp + b

		a, b = b, a + b
	}

	return b
}


func resolve02(n int) {
	// 动态规划解决爬楼梯问题
	// Dynamic programming to solve the stair climbing problem

	var res int = dynamicClimbing(n)

	fmt.Printf("动态规划解决: n = %d 阶楼梯，一共有 %d 种爬楼梯方法。", n, res)
}

func dynamicClimbing(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	var dp []int = make([]int, n, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}





func resolve01(n int) {
	// 递归解决爬楼梯问题
	// Recursive solution to the stair climbing problem

	var res int = recursiveClimbing(n)

	fmt.Printf("递归解决: n = %d 阶楼梯，一共有 %d 种爬楼梯方法。", n, res)

}

func recursiveClimbing(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return n
	} else {
		return recursiveClimbing(n -1) + recursiveClimbing(n - 2)
	}
}
