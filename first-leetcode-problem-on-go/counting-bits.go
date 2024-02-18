package main

import "fmt"

func main() {
	countBits(3)
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		num := i

		for num > 0 {
			count += num & i
			num >>= num
		}
		ans[i] = count
	}
	fmt.Println(ans)
	return ans
}
