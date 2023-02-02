package main

func SumArray(arr []int) (sum int) {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return
}
