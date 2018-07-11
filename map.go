// map : 添え字にキーをつかって、キーと値のペアで管理するデータ型（ハッシュ、連想配列）
package main

import "fmt"

func main() {
	// m := make(map[string]int) // キーがstring型、値がint型
	// m["foo"] = 100
	// m["baa"] = 200

	// 値を指定しながら宣言
	m := map[string]int{"foo": 300, "baa": 400}

	fmt.Println(m)
	fmt.Println(len(m)) //値の個数を調べる

	// 要素の削除
	delete(m, "foo")
	fmt.Println(m)

	// キーの存在を調べる
	v, ok := m["baa"]
	fmt.Println(v)
	fmt.Println(ok)

}
