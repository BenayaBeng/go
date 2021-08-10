package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("Nomor favorit saya adalah: ", rand.Intn(100))

	var i, j int = 1, 2
	var c = true
	var python = false
	var java = "no!"
	fmt.Println(i, j, c, python, java)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	l := 42
	m := 3.142
	n := 0.867 + 0.5i
	fmt.Printf("l = %T\n", l)
	fmt.Printf("m = %T\n", m)
	fmt.Printf("n = %T\n", n)

	q, r := 42, 2701
	p := &q
	fmt.Println(*p)
	*p = 21
	fmt.Println(q)

	p = &r
	*p = *p / 37
	fmt.Println(r)
}
