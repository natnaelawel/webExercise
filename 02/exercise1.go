package main

import "fmt"


func multTable(row , col int){
	for i:= 1 ; i<= row; i++{
		for j:=1 ; j<=col; j++{
			fmt.Print(i*j,"\t")
		}
		fmt.Println()
	}
}

func main() {
	multTable(12,12)
}
