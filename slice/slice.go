package main

import (
	"log"
	"fmt"
)

func main()  {
	RangeItemIsValueNotReference()

}

func RangeItemIsValueNotReference()  {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}
	// Iterate over each element and display the value and addresses.
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}
}

func AppendChangeUnderlayerArray()  {
	// Create a slice of integers. Contains a length and capacity of 5 elements.
	slice := []int{10, 20, 30, 40, 50}

	// Create a new slice. Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]
	log.Println("slice      = ", slice)
	log.Println("slice[1:3] = ", newSlice)
	newSlice = append(newSlice, 60)
	log.Println("After change slice, newSlice = append(newSlice, 60)")
	log.Println("slice      = ", slice)
	log.Println("slice[1:3] = ", newSlice)
}