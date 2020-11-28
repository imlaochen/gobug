package main

import "testing"
import "fmt"


/////////////////////////////////////////////////////////
// defer 按照定义顺序，后定义的先执行

// 输出：
// 222
// 111
func Test1(t *testing.T) {
	defer fmt.Println("111")
	defer fmt.Println("222")
}

/////////////////////////////////////////////////////////
// defer 函数定义的时候，参数已经被赋值了(参数被捕捉了)

// 输出：
// 第二个defer:  2
// 第一个defer:  1
func Test2(t *testing.T) {
	a := 1

	// 注意这里，a作为defer函数的参数。defer函数定义的时候，就把a的值确定了
	defer fmt.Println("第一个defer: ", a)


	defer func() {
		fmt.Println("第二个defer: ", a)
	}()

	a+=1
}


/////////////////////////////////////////////////////////
// defer 函数可以修改具名返回值，但是不会修改匿名返回值
// 输出
//   fHasName: 2
//   fNoName: 1
func fHasName() (val int) {
	val = 1
	defer func() {
		// 因为val是个具名返回值，所以这里对val的操作，会影响到返回值
		val += 1
	}()

	return val
}
func fNoName() int {
	val := 1
	defer func() {
		// 这里的操作则不会影响返回值
		val += 1
	}()

	return val
}
func Test3(t *testing.T)   {
	fmt.Println("fHasName:", fHasName())
	fmt.Println("fNoName:", fNoName())
}


////////////////////////////////////////////////////////
// 遇到panic的时候，当前协程调用栈已定义的defer将被执行.
// 如果当前协程没有任何一个defer函数内执行了recover，那么执行完所有的defer之后，将会触发panic退出进程

/* 输出
has: 333
err: haha
has: 222
has: 111
 */
func  TestHasRecover(t *testing.T)  {
	defer fmt.Println("has: 111")
	defer fmt.Println("has: 222")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	defer fmt.Println("has: 333")
	panic("haha")
	defer fmt.Println("has: 444")
}

/* 输出
has: 333
has: 222
has: 111

panic: haha [recovered]
	panic: haha

goroutine 19 [running]:
testing.tRunner.func1.1(0x95df20, 0x9b2110)
	C:/Go/src/testing/testing.go:1076 +0x310
testing.tRunner.func1(0xc000140480)
	C:/Go/src/testing/testing.go:1079 +0x43a
 */
func  TestNoRecover(t *testing.T)  {
	defer fmt.Println("has: 111")
	defer fmt.Println("has: 222")
	defer fmt.Println("has: 333")
	panic("haha")
	defer fmt.Println("has: 444")
}


/////////////////////////////////////////////////////

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