package main

import "testing"
import "fmt"

func sum(tip string, a int, b int) int {
	sum := a + b
	fmt.Printf("tip=%s a=%d b=%d sum=%d\n", tip, a, b, sum)
	return sum
}

func deferSum() {
	a := 1
	b := 2
	defer sum("1tip", a, sum("2tip", a, b))
	a = 3
	defer func(b int) {
		sum("3tip", a, sum("4tip", a, b))
	}(a)
	b = 4
	defer func() {
		sum("5tip", a, sum("6tip", a, b))
	}()
}

func TestDeferSum(t *testing.T) {
	deferSum()
}

//////////////////////////////////////////////////////////////////

func deferFunc1() int {
	start := 1
	defer func() {
		start += 1
		fmt.Println("defer1:", start)
	}()
	start += 1
	return start
}

func deferFunc2() int {
	start := 1
	defer func(start int) {
		start += 1
		fmt.Println("defer2:", start)
	}(start)
	start += 1
	return start
}
func deferFunc3() (start int) {
	start = 1
	defer func() {
		start += 1
		fmt.Println("defer3:", start)
	}()
	start += 1
	return start
}

func deferFunc4() (start int) {
	start = 1
	defer func(start int) {
		start += 1
		fmt.Println("defer4:", start)
	}(start)
	start += 1
	return start
}

func deferFunc5() (result int) {
	start := 1
	defer func(start int) {
		start += 1
		result += 1
		fmt.Println("defer5:", start)
	}(start)
	start += 2
	return start
}
func TestDeferFunc(t *testing.T) {
	fmt.Println(deferFunc1())
	fmt.Println(deferFunc2())
	fmt.Println(deferFunc3())
	fmt.Println(deferFunc4())
	fmt.Println(deferFunc5())
}
