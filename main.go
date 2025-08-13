package main

import(

	"fmt"
)

func main() {
	
	data, err := PartitionsDataTest()
	OldData := PartitionsData()	

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(OldData)


}
