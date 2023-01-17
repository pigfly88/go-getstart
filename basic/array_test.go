package basic

import "testing"

func TestVars(t *testing.T) {
	var a [3]int
	var b = [2]int{}
	c := [2]string{"a","b"}
	d := [...]string{"c","d","e"}
	t.Log(a, b, c, d)
}

func TestArray(t *testing.T) {
	arr1 := [3]int{1, 2, 3}
	for k, v := range arr1 {
		t.Log("k=", k, ",v=", v)
	}

	for k := range arr1 {
		t.Log("k=", k)
	}

	arr2 := [...]string{"a", "b", "c"}
	arr2[2] = "d"
	t.Log(arr2, " len=", len(arr2))

	reverse(&arr2)
	t.Log(arr2)
}

// Go语言中的函数参数传递，都是值传递，没有引用传递，slice在函数内部的改变会导致底层数组的改变是因为slice是一个结构体，而结构体里面的数组是一个指针。
func reverse(arr *[3]string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[2] = "fff"

}
