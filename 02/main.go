package main

import (
	"fmt"
)

func arrayfunc(){
	var nums [3]float64
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	for _, num := range nums{
		fmt.Println(num)
	}
	var slicenum = nums[:]
	for _, num := range slicenum{
		fmt.Println(num)
	}
}




type rectangle struct{
	width  int
	height int
}
func main(){
	// r1 := rectangle{width: 23,height:34}
	// r2 := rectangle{23,34}
	// r3 := new(rectangle)
	// r4 := rectangle{}

	// fmt.Println("area is ", r1.width * r1.height )

	// fmt.Println("area is ", r2.width * r2.height )
	// r3.width = 4
	// r3.height = 7
	// fmt.Println("area is ", r3.width * r3.height )

	// r3ptr := &r4
	// r3ptr.width = 6
	// fmt.Println("area is ", r3.width * r3.height )

	// r4.width = 4
	// r4.height = 5

	// fmt.Println("arae is ", r4.width * r4.height)

	arrayfunc()
}