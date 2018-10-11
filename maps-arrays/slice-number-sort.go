package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 5, 3, 6, 2, 10, 8}
	toBeSorted := sort.IntSlice(numbers)
	sort.Sort(toBeSorted)
	fmt.Println(toBeSorted)

	//reverse
	sort.Sort(sort.Reverse(toBeSorted))
	fmt.Println(toBeSorted)
}
