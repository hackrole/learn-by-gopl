package main

import "fmt"
import . "github.com/hackrole/gopl/ch1/add"
import "github.com/hackrole/gopl/ch1/add"

func main() {
	fmt.Println("vim-go")
	fmt.Println(Add(1, 2))
	fmt.Println(add.Add(1, 2))
}
