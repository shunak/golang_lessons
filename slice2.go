// make(),append(),copy()
package main

import "fmt"

func main() {

	// s := make([]int, 3) // [0,0,0]
	s := []int{1, 3, 5}
	// append : スライスに要素を加える
	s = append(s, 8, 2, 10)
	// copy：渡された値をコピーする
	t := make([]int, len(s))
	n := copy(t, s)
	fmt.Println(s) //[1,3,5,8,2,10]
	fmt.Println(t) //[1,3,5,8,2,10]
	fmt.Println(n) //6

}
