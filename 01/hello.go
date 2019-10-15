package main
import (
	"runtime"
	"fmt"
	// "strconv"
	// "os"
	// "math"
)
func pow(x, y int)(ans float64){
	ans = 1
	for i:= 1; i<=y; i++{
		ans *= float64(x)
	}
	return ans
}
func swap( x, y float64) (float64, float64){
	return y, x
}
func sumofN(num int)(float64){
	var sum float64 = 0
	for i := 1 ; i<= num; i++{
		sum += float64(i)
	}
	return sum
}
func average(num int)(float64){
	return sumofN(num)/ float64(num)
}
// func compare(x, y float64)(string){
// 	if (x > y) {
// 		return str
// 			}else if x < y{
// 		return string(y)
// 	}else{
// 		fmt.Println("Both are Equal")
// 		return float64(nil)
// 	}
// }


func pointerfunc(){
	var ptr *int
	var nptr = new(int)
	*nptr = 3
	// we can also you the new(int) inorder to create a pointer
	var x = 10
	ptr = &x
	y := *ptr
	fmt.Println(ptr)
	fmt.Println(x)
	fmt.Printf("%T\n", ptr)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Println(*nptr)
	fmt.Println(&nptr)
}

func deferfunc(){
	fmt.Println("defer function is started")

	defer fmt.Println("counting end")
	for i:= 0 ; i<= 10 ;i++{
		defer fmt.Println(i)
	}
	defer fmt.Println("counting started")
}

func switchfunc(){
	switch os := runtime.GOOS; os {	
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
		fallthrough
	case "darwin":
		fmt.Println("OS .x")
	default:
		fmt.Println(os)
}
}

func main(){
	pointerfunc()
	// deferfunc()
	// switchfunc()
	// fmt.Println(compare(6, 6))
	// fmt.Println("sum of the first 100 number is ",sumofN(100))
	// fmt.Println("average of the first 100 ",average(100))
	// fmt.Println(pow(2, 10))
	// x, y := swap(4,9)
	// fmt.Println(x, y)
	// fmt.Println("hello world")
	// fmt.Println("my favourite number is ", rand.Intn(10))

	// for i := 0; i < 10; i++ {
	// 	fmt.Print(rand.Intn(10),", ")
	// }
	// var x, y int = 3, 6
	// var f float32 = float32(math.Sqrt(float32(x * x + y * y)))
	

	// fmt.Println(num)
}