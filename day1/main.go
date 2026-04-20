package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World")

	name := "World"

	// コマンドライン引数があれば使う。
	// Argsの[0]はプログラム名なので、引数を使うときは[1]以降を使う
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 環境変数が未設定の場合は、空文字
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello"
	}

	fmt.Printf("%s, %s!\n", greeting, name)
}