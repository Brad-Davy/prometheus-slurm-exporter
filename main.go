package main

import "fmt"

func main() {
		result,err := PartitionsDataTest()
	if err != nil{
		fmt.Println(err)
	}

	newresult := PartitionsData()
    fmt.Println("Old Output")
    fmt.Println(string(newresult))
	fmt.Println("New Output")
	fmt.Println(result)
}
