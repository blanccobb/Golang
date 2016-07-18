package main

import (
	"fmt"
	"math"
)

var x, y, z int
var c, python, java bool

func main(){
	fmt.Println(hanoi(3,"a","b","c"))
	//use_slice()
	//iterate_slice()
	//a, b := swap("hello","world")
	//fmt.Println(a,b)
	//fmt.Println(split(17))
	//fmt.Println(x, y, z, c, python, java)
	//shortInit()
	
	//2016-06-21
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	
	//2016-07-12
	fmt.Println("Newton's method : ",  Sqrt(-25))
	fmt.Println("math.Sqrt method : ", math.Sqrt(25))
	
	//2016-07-18
	fmt.Println(Vertex{1,2})
	
	v := Vertex{1,2}				//X = 1 , Y = 2
	v.X = 4
	fmt.Println(v.X, " : ", v.Y)	//X = 4 , Y = 2
	pointerCopy(v)
	exam_Literals()
	
	exam_newFunc()
	test_slice()
	slicing_slices()
	
	a := make([]int, 5)
	print_slice("a", a)
	b := make([]int, 0, 5)
	print_slice("b", b)
	c := b[:2]
	print_slice("c", c)
	d := c[2:5]
	print_slice("d", d)
}