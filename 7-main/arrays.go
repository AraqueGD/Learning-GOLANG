package main

import "fmt"

func main() {
	list := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < len(list); i++ {
		fmt.Println(list[i])
	}
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice2()
	slice3()
	slice4()

}
func slice2() {
	elments := [5]int{4, 5, 6, 7, 8}
	portion := elments[3:]
	fmt.Println(portion)
}

func slice3() {
	elements := make([]int, 5, 20)
	for i := 0; i < 100; i++ {
		elements = append(elements, i)
	}
	fmt.Printf("Len: %d, Cap: %d", len(elements), cap(elements))
}

func slice4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLen: %d, Cap: %d", len(nums), cap(nums))
}
