// Package solution は「初めてのGo 第2版」の10章に関するパッケージです。
package solution

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add は2つの整数を足し算して、その結果を返します。
// 足し算については[mathsisfun]のサイトを参考にしてください。
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
