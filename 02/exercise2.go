package main
import "fmt"
func diamond(){
var diagonal int = 10
var space int = 5
var num = 0
for i:=0; i<= diagonal; i++{
	for j:= 1 ; j<= space; j++{
		fmt.Print(" ")
	}
	for k:=1 ; k <= (2* num )+1; k++{
		fmt.Print("*")
	}
	fmt.Println()
	
	if i >= diagonal/2 {
		space++
		num--
	}else{
		space--
		num++
	}
}
}
func main() {
       diamond()
}