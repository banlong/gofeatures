package main

import (
	"sync"
	"strconv"
	"fmt"
)

// Fibonacci sequence, and characterized by the fact that every number after the first two is
// the sum of the two preceding ones:[1][2]
// The Fibonacci numbers for n=1, 2, ... are 1, 1, 2, 3, 5, 8, 13, 21, ...
func main()  {
	var n int64 = 10
	remain := n%10
	var order string
	switch remain {
	case 1:
		order = "st"
		break
	case 2:
		order = "nd"
		break
	case 3:
		order = "rd"
		break
	default:
		order = "th"
		break
	}

	f := Fib(n)
	fmt.Printf("The " + strconv.FormatInt(n ,10) +  order + " Fibonacci number is: \n")
	fmt.Printf("%v \n", f)
}

//fmt.Stringer() returns a string representation of its receiver
func Fib(n int64) fmt.Stringer{
	var promise FutureInt64
	promise.ch = make(chan int64)

	go func() {
		_, f := fib(n)

		//the final value is sent to a channel
		promise.ch <- f
	}()

	//a string is returned from the FutureInt64 type
	return &promise
}

//Recursive to get fib value
func fib(n int64) (int64, int64) {
	if n < 2 {
		fmt.Printf("n=%v, f1=%v, f2=%v \n", n, 1, 1)
		return 1, 1
	}
	f1, f2 := fib(n-1)
	fmt.Printf("n=%v, f1=%v, f2=%v \n", n, f1, f2)
	return f2, f1+f2
}

type FutureInt64 struct {
	ch chan int64
	v int64
	collect sync.Once
	//Once is an object that will perform exactly one action.
}

func (f FutureInt64) String() string {
	f.collect.Do(func() {
		f.v = <- f.ch
	})
	return  strconv.FormatInt(f.v, 10)
}
