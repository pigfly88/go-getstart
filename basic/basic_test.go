package basic_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

// 全局变量
var (
	globalVar = 666
)

// 常量
const (
	SUCCESS = iota // 0
	ERROR          // 1
	UNKNOWN        // 2
)

// 常量 iota结合位运算
const (
	READABLE   = 1 << iota // 1 << 0 is 001
	WRITABLE               // 1 << 1 is 010
	EXECUTABLE             // 1 << 2 is 100
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello World")
}

func TestVars(t *testing.T) {

	var x1 *int
	t.Logf("%T, %v", x1, x1)

	var x2 = new(int)
	t.Logf("%T, %v, %v", x2, x2, *x2)

	c := make([]int, 1)
	t.Logf("%T, %v", c, c)
	// 1.只定义变量类型 不赋值
	var i1 int

	// 2.定义变量类型并赋值
	var i2 int = 1

	// 省略类型，go编译器会自动推断类型
	var i3 = 2

	// 4.更简洁的方式
	i4 := 3

	// 多变量声明
	i5, i6 := 4, 5

	t.Log(i1, i2, i3, i4, i5, i6, globalVar, SUCCESS, ERROR, UNKNOWN)

	guestPermission := READABLE
	userPermission := READABLE | WRITABLE
	adminPermission := READABLE | WRITABLE | EXECUTABLE
	t.Log(guestPermission&READABLE, guestPermission&WRITABLE, userPermission&WRITABLE, userPermission&EXECUTABLE, adminPermission&EXECUTABLE)

	// 作用域
	x := 1
	t.Log(x) //prints 1
	{
		t.Log(x) //prints 1
		x = 3
		x := 2   // 不会影响到外部x变量的值
		t.Log(x) //prints 2
		//x = 5        // 不会影响到外部x变量值
	}
	t.Log(x) //prints 3

	// 字符串不允许使用nil值
	//var str1 string = nil

	str2 := "abc"
	str2 = "def"
	t.Log(str2)
}

func TestArgs(t *testing.T) {
	for _, arg := range os.Args {
		t.Log(arg)
	}
}

func TestFunc(t *testing.T) {
	a, b := 1, 2
	swap(&a, &b)
	t.Log(a, b)
}

func init() {
	fmt.Println("init()")
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

func TestMultiReturn(t *testing.T) {
	t.Log(multiReturn1(1, 2))
	t.Log(multiReturn2(1, 2))
}

func multiReturn1(a int, b int) (int, int) {
	return a + 10, a + b
}

func multiReturn2(a int, b int) (x int, y int) {
	x = a + 20
	y = a + b
	return
}

func TestSlice(t *testing.T) {
	// 定义
	var slice1 []int
	var slice2 = make([]int, 6)
	slice31 := []int{1,2,3}
	t.Log(slice1, len(slice1), slice2, len(slice2), slice31)

	slice1 = append(slice1, 1)
	slice2[0] = 9
	slice2 = append(slice2, 8)

	t.Log(slice1, len(slice1), cap(slice1), slice2, len(slice2), cap(slice2))

	slice3 := slice2[:3]
	t.Log(slice3, len(slice3), cap(slice3), slice3[:12])

	changeValue(slice3)
	t.Log(slice3, slice2)

	slice4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // len=10 cap=10
	var slice7 = make([]int, 10)
	copy(slice7, slice4)    // copy不共用底层数组
	slice5 := slice4[2:5]   // [2,3,4], len=3, cap=8
	slice6 := slice5[2:6:7] // [4,5,6,7], len=4, cap=5

	slice6 = append(slice6, 100) // [4,5,6,7,100], len=5, cap=5
	slice6 = append(slice6, 200) // 容量不足，创建新的底层数组 [4,5,6,7,100,200], len=6, cap=10

	slice4[2] = 20
	t.Log(slice4, slice5, slice6, slice7)

	aa := []int{1, 2, 3, 4, 5, 6}
	bb := aa[2:5] // [3,4,5],len=3,cap=4
	cc := append(bb, 11)
	//bb = append(bb, 11) // [3,4,5,11],len=4,cap=4
	//bb = append(bb, 12) // 容量不足，创建新的底层数组[3,4,5,11,12],len=5,cap=8
	t.Log(aa, bb, cc)

	a := []int{1} // len=1,cap=1
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a = append(a, 2) // 扩容,len=2,cap=2
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a = append(a, 3) // 扩容,len=3,cap=4
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	b := append(a, 4) // 由于 a 的底层数组仍然有空间，因此并不会扩容。这样，底层数组就变成了 `[1,2,3,4]`。注意，此时 a = `[1,2,3]`，容量为4；b = `[1,2,3,4]`，容量为4。这里 a 不变
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&b)))
	c := append(a, 5) // 由于 s 的len=3,cap=4,所以还是不需要扩容，因此直接底层数组索引为3的地方填上5，底层数组变成[1,2,3,5],注意这里a不变
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&a)))
	t.Log((*reflect.SliceHeader)(unsafe.Pointer(&c)))
	t.Log(a, b, c)
	// 这里要注意的是，append函数执行完后，返回的是一个全新的 slice，并且对传入的 slice 并不影响。append的元素取决于传入的slice的len，而不是简单在底层数组后面追加元素。
}

// Go语言中的函数参数传递，都是值传递，没有引用传递，slice在函数内部的改变会导致底层数组的改变是因为slice是一个结构体，而结构体里面的数组是一个指针。
func changeValue(s []int) {
	//s[1] = 8
	s = append(s, 99)
}

func TestMap(t *testing.T) {
	map1 := map[string]string{
		"one": "1",
		"two": "2",
	}

	if var1, ok := map1["three"]; !ok { // 判断是否存在，x[key]始终有返回值
		t.Log("key not exists")
	} else {
		t.Log(var1)
	}

	map2 := make(map[string]int)
	map2["golang"] = 1
	map2["php"] = 2
	t.Log(map2)

	changeValue2(map2)
	t.Log(map2)

	s := []int{111, 999}
	changeValue2(s)
	t.Log(s)

	var res []int = twoSum([]int{5, 1, 9, 4, 2, 8}, 12)
	t.Log(res)
}

func changeValue2(s interface{}) {
	switch s := s.(type) {
	case map[string]int:
		s["java"] = 3
	case []int:
		s[0] = 222
	default:

	}

}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
