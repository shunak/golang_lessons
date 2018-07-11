// 構造体:複数の値を意味のあるまとまりとして新しい型を定義することができる
// メソッド（データ型に紐づいた関数）
package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

// (u user)のようにレシーバをセットする（値渡し）
func (u user) show() {
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

// レシーバの値を(u *user)のように参照渡しにするとuの値をアドレスから
// 参照するようになる
func (u *user) hit() {
	u.score++
}

func main() {
	u := user{name: "foo", score: 20}
	// fmt.Println()
	u.hit()
	u.show()
}
