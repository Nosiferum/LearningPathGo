package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	i, j := 42, 350

	var k *int
	k = &j
	fmt.Println(&j)
	fmt.Println(k)
	fmt.Println(&k)
	fmt.Println(*k)

	p := &i
	fmt.Println(*p)
	*p = 25
	fmt.Println(i)
	p = &j
	*p = *p / 10
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})
	v := Vertex{2, 3}
	v.X = 5
	fmt.Println(v.X)

	p2 := &v
	p2.X = 50
	fmt.Println(v)

	p3 := &Vertex{5, 10}
	fmt.Println(p3)
	fmt.Println(*p3)

	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slice
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	structs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structs)

	slice := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice)

	// Slice the slice to give it zero length.
	slice = slice[:0]
	printSlice(slice)

	// Extend its length.
	slice = slice[:4]
	printSlice(slice)

	// Drop its first two values.
	slice = slice[2:]
	printSlice(slice)

	dynamicArr := make([]int, 5)
	printSlice(dynamicArr)

	dynamicArr2 := make([]int, 0, 5)
	printSlice(dynamicArr2)

	dynamicArr3 := dynamicArr2[:2]
	printSlice(dynamicArr3)

	dynamicArr4 := dynamicArr3[2:5]
	printSlice(dynamicArr4)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var slice2 []int
	printSlice(slice2)

	// append works on nil slices.
	slice2 = append(slice2, 0)
	printSlice(slice2)

	// The slice grows as needed.
	slice2 = append(slice2, 1)
	printSlice(slice2)

	// We can add more than one element at a time.
	slice2 = append(slice2, 2, 3, 4)
	printSlice(slice2)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)

	for i := range pow2 {
		pow2[i] = i + 1
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	var m map[string]Vertex
	m = make(map[string]Vertex)

	m["test"] = Vertex{1, 2}
	fmt.Println(m["test"])

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": Vertex{
			37, -122,
		},
	}
	fmt.Println(m2)

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	value, ok := m3["Answer"]
	fmt.Println("The value:", value, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
