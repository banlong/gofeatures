package main

import "fmt"

//An interface can be satisfied (or implemented) by an arbitrary number of types.
// Here, the interface Men is implemented by both Student, Human, and Employee.

// Also, a type can implement an arbitrary number of interfaces, here, Student implements (or satisfies)
// both Men and YoungChap interfaces, and Employee satisfies Men and ElderlyGent.

//And finally, every type implements the empty interface and that contains, you guessed it: no methods.
// We declare it as interface{} (Again, notice that there are no methods in it.)

// If we declare m of interface Men, it may store a value of type Student or Employee, or even... (gasp) Human!
// This is because of the fact that these type all implement methods specified by the Men interface.

// If m (type Men) can store values of these different types, we can easily declare a slice of type Men that will
// contain heterogeneous values(Student, Employee, or Human). This was not even possible with slices of classical types!


func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	//a variable of the interface type Men
	var i Men

	//i can store a Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i can store an Employee too
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//a slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//These elements are of different types that satisfy the Men interface
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}

type Human struct {
	name string
	age int
	phone string
}

// A human likes to stay... err... *say* hi
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// A human can sing a song, preferrably to a familiar tune!
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// A Human man likes to guzzle his beer!
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}



type Employee struct {
	Human //an anonymous field of type Human
	company string
	money float32
}

// Employee's method for saying hi overrides a normal Human's one
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// An Employee spends some of his salary
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}


type Student struct {
	Human //an anonymous field of type Human
	school string
	loan float32
}

// A Student borrows some money
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}



// INTERFACES
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
