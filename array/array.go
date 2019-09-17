package main

import "fmt"

func main() {
	var i, j, k int
	var q string
	var arr = [...]int{1, 2, 3}
	i = 2
	j = 3
	k = 4
	q = "quyen"
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%s is\n", q)

	fmt.Println("Array is")
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}

	var arr01 = [...]int{8, 9, 10, 4: 12}
	fmt.Println("Array01 is")
	for i, v := range arr01 {
		fmt.Printf("b[%d] is %d\n", i, v)
	}

	var arr02 = [...]int{55, 7, 15, 5: 12}
	fmt.Println("Array02 is")
	for i = 0; i < len(arr01); i++ {
		fmt.Printf("b[%d] is %d\n", i, arr02[i])
	}
}
