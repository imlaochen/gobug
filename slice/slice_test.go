package main

import (
	"fmt"
	"testing"
)

/*
输出结果：
array1: [1 2 3] array2 [100 2 3]
slice1: [100 2 3] slice2 [100 2 3]
 */
func TestSlice1(t *testing.T) {
	// 数组是赋值传递
	array1 := [3]int{1,2,3}

	// 这里array1复制了一份，array2和array1已经不是同一份数据了
	// 所以对数组array1，array2的修改是互不影响的
	array2 := array1
	array2[0] = 100
	fmt.Println("array1:", array1, "array2", array2)



	// 切片是引用传递
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1, 2, 3)

	// slice2和slice1引用的同一份数据，所以slice2对已有元素的修改，会影响到slice1
	// 这里有个细节需要注意，往下面TestSlice2继续看
	slice2 := slice1
	slice2[0] = 100
	fmt.Println("slice1:", slice1, "slice2", slice2)
}

/*
输出结果
slice1: [1 2 3] slice2 [1 2 3 4]
slice1: [1 2 3] slice2 [1 2 3 4 5]
slice1: [1 2 3] slice2 [100 2 3 4 5]
 */
func TestSlice2(t *testing.T) {
	// 切片是引用传递
	slice1 := make([]int, 0, 4) // 容量是4
	slice1 = append(slice1, 1, 2, 3)

	// slice2和slice1引用的同一份数据，所以slice2对元素的修改，会影响到slice1
	// 特别要注意的是，这里说的引用同一份数据，实际上是指的slice1和slice2内部的ptr指向了同一个数组
	// 但是slice1和slice2结构中的len和cap是复制传递的
	slice2 := slice1
	// 添加第4个元素，容量足够，不会触发内部ptr数组重新分配
	slice2 = append(slice2, 4)
	fmt.Println("slice1:", slice1, "slice2", slice2)

	// 添加第5个元素，容量不足，slice2.ptr重新分配内存，此时slice2.ptr和slice1.ptr已经不是同一份内存了
	slice2 = append(slice2, 5)
	fmt.Println("slice1:", slice1, "slice2", slice2)

	// 所以这里slice2对已有元素的修改，不会影响到slice1了
	slice2[0] = 100
	fmt.Println("slice1:", slice1, "slice2", slice2)
}

/*
go/src/runtime/slice.go slice结构定义：
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
 */

/*
type Slice struct {
	Array []interface{}
	Len int
	Cap int
}

func fake()  {
	slice1 := &Slice{
		Array: 指向一块连续内存的定长数组,
		Len: 0,
		Cap: 4
	}

	slice2 := slice1 // 这一行相当于如下：

	slice2 := &Slice{}
	slice2.Array = slice1.Array // 指向同一块内存，所以说切片是引用传递
	slice2.Len = slice1.Len // len字段被复制了
	slice2.Cap = slice2.Cap // cap字段被复制了

	// 后续如果触发了任一slice的array重新分配内存，另一个slice都是不知道的
	// 已经修改任一slice的len和cap，另一个slice也是不知道的
}
 */