package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	answer := rand.IntN(100) + 1
	attempts := 0
	var start time.Time = time.Now()

	fmt.Println("===== 数当てゲーム =====")
	fmt.Println("1〜100の数を当ててください！")
	fmt.Println()

	for {
		fmt.Print("予想: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(" -> 数値を入力してください")
			continue
		}

		attempts++

		if guess < answer {
			fmt.Println(" -> もっと大きい！")
		} else if guess > answer {
			fmt.Println(" -> もっと小さい！")
		} else {
			// time.Sinceでstartから経過した時間を取得できる
			var elapsed time.Duration = time.Since(start)
			fmt.Println()
			fmt.Printf("正解！ %d でした！\n", answer)
			fmt.Printf("試行回数: %d回\n", attempts)
			fmt.Printf("所要時間: %.1f秒\n", elapsed.Seconds())
			break
		}
	}
}