package _defer

import "testing"

/**
函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。如果需要在defer里面修改返回值并生效的话，需要在函数定义里写明返回的变量名。
如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作，所以在defer里面修改不会生效。
而有名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，在执行return的时候会先执行返回值保存操作，
而后续的defer函数会改变这个返回值(虽然defer是在return之后执行的，但是由于使用的函数定义的变量，所以执行defer操作后对该变量的修改会影响到return的值
*/
func TestReturn(t *testing.T) {
	t.Log(returnArg0())  // 5
	t.Log(returnArg1(1)) // 2
	t.Log(returnArg2(1)) // 0
	t.Log(returnArg3(1)) // 2
	t.Log(returnArg4())  // 5
	t.Log(returnArg5())  // 10
}

func returnArg0() int {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func returnArg1(x int) (y int) {
	y = x + 1
	return
}

func returnArg2(x int) (y int) {
	y = x + 1
	return 0
}

func returnArg3(x int) (y int) {
	y = 6
	defer func() {
		y = x + 1
	}()
	return 0 // 先赋值：y = 0，然后执行defer，因为defer里面对y做了修改，所以返回值 y = 1 + 1 = 2
}

func returnArg4() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t // 先赋值：r = t，然后执行defer，因为defer里面是对t值的操作不会影响到r，所以r还是5
}

func returnArg5() (r int) {
	t := 5
	defer func() {
		r = t + 5
	}()
	return t // 先赋值：r = t，然后执行defer，因为defer里对r做了修改，所以会影响返回值
}
