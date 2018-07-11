// range：slice,map等とあわせてつかえる
package main

import (
	"fmt"
)

func main() {
	// スライスの作成
	// s := []int{2, 3, 8}
	// rangeでsの要素をすべてを取り出しながら、i,vに値を格納する
	// for i, v := range s {
	// 	fmt.Println(i, v)
	// }
	// ブランク修飾子_：不要な要素を除くためのもの
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

	// mapの作成
	m := map[string]int{"foo": 100, "baa": 200}
	// mapの場合にはkeyとvalueがかえってくる
	for k, v := range m {
		fmt.Println(k, v)
	}

}
