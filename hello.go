package main

import (
	"fmt"
)

func binarySearch(arr [5]int, l int, r int, x int) int {

	if r >= l {
		mid := (r + l) / 2

		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		} else {
			return binarySearch(arr, mid+1, r, x)
		}
	}
	return -1
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("enter the number")
	var x int
	fmt.Scanln(&x)
	ans := binarySearch(a, 0, 5, x)
	fmt.Println(ans)
}
