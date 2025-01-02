package main

func fibonacci(n int) int {

	sliceFib := []int{0, 1}
	appNum := 0

	for i := 2; i <= n; i++ {
		appNum = sliceFib[i-1] + sliceFib[i-2]
		sliceFib = append(sliceFib, appNum)
	}
	return sliceFib[n]
}
