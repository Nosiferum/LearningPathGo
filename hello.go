package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
)

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Pi  = 3.14
	aha = "aha"
)

func main() {
	fmt.Println("Hello, 世界")

	fmt.Println(time.Now())
	fmt.Println(rand.Int())

	fmt.Println(add(10, 50), "is", add(29, 32))

	var variable string
	variable = "a"
	fmt.Println(variable)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	variable2 := "b"
	fmt.Println(variable2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	i := 42
	f := float64(i)
	u := uint(i)

	fmt.Println(i, f, u)

	const World = "世界"
	fmt.Println(World)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(25))

}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
