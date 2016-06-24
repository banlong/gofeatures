package main
import "fmt"

//Recursive Max
func Max(slice []int) int{
	if len(slice) == 1 {
		return slice[0] //there's only one element in the slice, return it!
	}

	middle := len(slice)/2 //the middle index of the slice
	//find out the Max of each sub-slice
	m1 := Max(slice[:middle])
	m2 := Max(slice[middle:])
	//compare the Max of two sub-slices and return the bigger one.
	if m1 > m2 {
		return m1
	}
	return m2
}

func main(){
	s := []int {1, 2, 3, 4, 6, 8}
	fmt.Println("Max(s) = ", Max(s))
}