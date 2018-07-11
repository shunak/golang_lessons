// インターフェース
// メソッドの一覧を定義したデータ型
// それらのメソッドをある構造体が実装していればその構造体をそのインターフェース型として扱っても良いという特徴がある

package main

import (
	"fmt"
)

func main() {
	u := user{name: "baa", score: 20}
	fmt.Println(u)
}
