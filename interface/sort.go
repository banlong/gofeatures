package main
import (
	"fmt"
	"strconv"
	"sort"
)

//The sort package provides functions to sort slices of type int, float, string and wild and wonderous things.
// In fact, the sort package defines a interface simply called "Interface" that contains three methods:
/* type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
  } */

//From the sort Interface doc:
// “A type, typically a collection, that satisfies sort. Interface can be sorted by the routines in this package.
// The methods require that the elements of the collection be enumerated by an integer index.
// -->So all we need in order to sort slices of any type is to implement these three methods!
// Let’s give it a try with a slice of Human that we want to sort based on their ages.


type Human struct {
	name string
	age int
	phone string
}

func (h Human) String() string {
	return "(name: " + h.name + " - age: "+strconv.Itoa(h.age)+ " years)"
}

type HumanGroup []Human //HumanGroup is a type of slices that contain Humans

func (g HumanGroup) Len() int {
	return len(g)
}

func (g HumanGroup) Less(i, j int) bool {
	if g[i].age < g[j].age {
		return true
	}
	return false
}

func (g HumanGroup) Swap(i, j int){
	g[i], g[j] = g[j], g[i]
}

func main(){
	group := HumanGroup{
		Human{name:"Bart", age:24},
		Human{name:"Bob", age:23},
		Human{name:"Gertrude", age:104},
		Human{name:"Paul", age:44},
		Human{name:"Sam", age:34},
		Human{name:"Jack", age:54},
		Human{name:"Martha", age:74},
		Human{name:"Leo", age:4},
	}

	//Let's print this group as it is
	fmt.Println("The unsorted group is:")
	for _, v := range group{
		fmt.Println(v)
	}

	//Now let's sort it using the sort.Sort function
	sort.Sort(group)

	//Print the sorted group
	fmt.Println("\nThe sorted group is:")
	for _, v := range group{
		fmt.Println(v)
	}
}
