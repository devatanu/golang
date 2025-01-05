package main

import "fmt"

// setArray modifies the array at the specified index and returns a boolean indicating success.
func setArray(arr []int, index int, value int) bool {
	if index >= 0 && index < len(arr) {
		arr[index] = value
		return true
	} else {
		fmt.Println("Error: Index out of range")
		return false
	}
}

func main() {
	var a [10]int
	fmt.Println("emp: ", a)
	setArray(a[:], 0, 100)
	fmt.Println("set: ", a)
}
