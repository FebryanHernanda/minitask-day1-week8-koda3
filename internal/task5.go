package internal

import "fmt"

func Task5() {
	var sliceTemp = []int{50, 75, 66, 20, 32, 90}

	fmt.Println("Before : ", sliceTemp)
	sliceTemp = append(sliceTemp[:3], append([]int{88}, sliceTemp[3:]...)...)
	fmt.Println("After : ", sliceTemp)
}
