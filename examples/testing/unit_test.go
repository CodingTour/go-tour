package testing

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	t.Log(Swap(0, 1))
}

func Example() {
	fmt.Println(Swap(0, 1))
	// Output: 1 0
}

/*
在Java中交换两个变量需要一个临时变量:
int temp = x
x = y
y = temp
*/
func Swap(x int, y int) (int, int) {
	x, y = y, x
	return x, y
}



