package main
import (
	"fmt"
	"math/rand"
	)
	
func contains(slice []int, element int) bool {
    for _, value := range slice {
        if value == element {
            return true
        }
    }
    return false
}
func guessGame(){
	randomNum := rand.Intn(100)
	var guessedNum int
	var tryNum int = 4
	var guessedNumSlice[] int

	for i := tryNum; i >= 0; i--{
		fmt.Println("Enter Your Guess")
		fmt.Scan(&guessedNum)
		if contains(guessedNumSlice, guessedNum ){
			fmt.Println("You Already Guessed")
			i++
			continue
		}else{
		guessedNumSlice = append(guessedNumSlice, guessedNum )
		if guessedNum < randomNum {
			fmt.Println("You Guessed Low")
			fmt.Printf("You have %d more try \n", i)
		} else if guessedNum >  randomNum {
			fmt.Println("You Guessed High")
			fmt.Printf("You have %d more try \n", i)

		} else {
			fmt.Println("You Guessed Right")
	}
		
	}
}
}
func main(){
       guessGame()
}