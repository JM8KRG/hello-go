package main

import "fmt"

func main() {
	var gafa[5]string = [5]string{"Google", "Apple", "Facebook", "Amazon"}
	
	for i := 0; i < len(gafa); i++ {
		fmt.Println(gafa[i])
	}
}
