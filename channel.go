// goroutine : 並行処理
// チャンネル: データの受け渡井をするパイプ

package main

import (
	"fmt"
	"time"
)

// task1の方でchannelに対してデータを渡して、
// mainの方でそのデータを取り出す
func task1(result chan string) {
	time.Sleep(time.Second * 2)
	// fmt.Println("task1 finised!")
	// チャンネルresultに対して、データを渡している
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finised!")

}

// goというprefixでを付すことで、並列計算が可能になる
func main() {
	// チャンネルの作成
	result := make(chan string)

	// 同時に走らせたい処理がある場合には go を付けて goroutine にすることができる
	go task1(result)
	go task2()

	fmt.Println(<-result)

	// goroutine が終わる前に main 関数自体が終わってしまうと何もおきないので、下記を付与
	time.Sleep(time.Second * 3)

}
