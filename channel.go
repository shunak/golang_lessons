// goroutine : 並行処理
// チャンネル: データの受け渡井をするパイプ

package main

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	time.Sleep(time.Second * 2)
	// fmt.Println("task1 finised!")
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finised!")

}

// go のprefixでを付すことで、並列計算が可能になる
func main() {
	result := make(chan string)

	// 同時に走らせたい処理がある場合には go を付けて goroutine にすることができる
	go task1(result)
	go task2()

	fmt.Println(<-result)

	// goroutine が終わる前に main 関数自体が終わってしまうと何もおきないので、下記を付与
	time.Sleep(time.Second * 3)

}
