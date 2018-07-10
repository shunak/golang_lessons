// ポインタ
// アドレスを指し示す変数
// ポインタの演算はできない
package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a //&a = aのアドレス
	// paの領域にあるデータの値 = *pa
	fmt.Println(pa)  //アドレスの表示(16進数)
	fmt.Println(*pa) //5

}
