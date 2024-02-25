package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

func main() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex2{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	var abser Abser

	abser = f
	abser = &v

	// In the following line, v is a Vertex2 (not *Vertex2)
	// and does NOT implement Abser.
	// abser = v

	fmt.Println(abser.Abs2())

	var myInterface MyInterface = MyType{"hello"}
	myInterface.Method()

	// nil type
	var myType2 *MyType2
	myInterface = myType2
	describe(myInterface)
	myInterface.Method()

	myInterface = &MyType2{"hello"}
	describe(myInterface)
	myInterface.Method()

	myInterface = MyFloat2(math.Pi)
	describe(myInterface)
	myInterface.Method()

	var myInterface2 interface{}
	describe2(myInterface2)

	myInterface2 = 42
	describe2(myInterface2)

	myInterface2 = "hello"
	describe2(myInterface2)

	myString := myInterface2.(string)
	fmt.Println(myString)

	myString, ok := myInterface2.(string)
	fmt.Println(myString, ok)

	myFloat, ok2 := myInterface2.(float64)
	fmt.Println(myFloat, ok2)

	//panic
	/*myFloat2 := myInterface2.(float64)
	fmt.Println(myFloat2)*/

	do(21)
	do("hello")
	do(true)

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}

type Vertex2 struct {
	X, Y float64
}

// Abs pass by value
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale pass by reference
func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs2() float64
}

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyInterface interface {
	Method()
}

type MyType struct {
	myString string
}

type MyType2 struct {
	myString string
}

type MyFloat2 float64

func (myType MyType) Method() {
	fmt.Println(myType.myString)
}

func (myType2 *MyType2) Method() {
	if myType2 == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(myType2.myString)
}

func (myFloat MyFloat2) Method() {
	fmt.Println(myFloat)

}

func describe(myInterface MyInterface) {
	fmt.Printf("(%v, %T)\n", myInterface, myInterface)
}

func describe2(myInterface interface{}) {
	fmt.Printf("(%v, %T)\n", myInterface, myInterface)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (person Person) String() string {
	return fmt.Sprintf("%v (%v years)", person.Name, person.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
