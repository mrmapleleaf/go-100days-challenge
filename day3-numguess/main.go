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

	for {
	
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("===== 数当てゲーム =====")
		fmt.Println("難易度を選んでください:")
		fmt.Println("  1: かんたん (1〜50)")
		fmt.Println("  2: ふつう  (1〜100)")
		fmt.Println("  3: むずい  (1〜1000)")
		fmt.Print("選択: ")

		scanner.Scan()
		maxNum, maxAttempts := chooseDificulty(strings.TrimSpace(scanner.Text()))

		answer := rand.IntN(maxNum) + 1
		attempts := 0
		var start time.Time = time.Now()

		fmt.Println()
		fmt.Printf("1〜%dの数を当ててください（最大%d回）\n", maxNum, maxAttempts)
		fmt.Println()

		for {
			remaining := maxAttempts - attempts
			fmt.Printf("予想（残り%d回）: ", remaining)
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			guess, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(" -> 数値を入力してください")
				continue
			}

			if guess < 1 || guess > maxNum {
				fmt.Printf(" -> 1〜%dの範囲で入力してください\n", maxNum)
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
				printRank(attempts, maxAttempts)
				break
			}

			if attempts >= maxAttempts {
				fmt.Println()
				fmt.Printf("ゲームオーバー！ 正解は %d でした\n", answer)
				break
			}
		}

		// リプレイ確認
		fmt.Println()
		fmt.Print("もう一度遊ぶ？ (y/n): ")
		scanner.Scan()
		if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
			fmt.Println("ありがとうございました！")
			break
		}
		fmt.Println()
	}
}

// Goは複数の返り値を渡せるので、わかりやすくなるように返り値に名前をつけられる
// 受け取る側は同じ名前で受ける
func chooseDificulty(choice string) (maxNum int, maxAttempts int) {
	switch choice {
	case "1":
		return 50, 8
	case "3":
		return 1000, 12
	default:
		return 100, 10
	}
}

func printRank(attempts, maxAttempt int) {
	ratio := float64(attempts) / float64(maxAttempt)
	// switchでcase文に直接条件式を書くパターン
	switch {
		case attempts == 1:
			fmt.Println("ランク: S（神）")
		case ratio <= 0.3:
			fmt.Println("ランク: A（すごい！）")
		case ratio <= 0.6:
			fmt.Println("ランク: B（いいね！）")
		case ratio <= 0.8:
			fmt.Println("ランク: C（まずまず）")
		default:
			fmt.Println("ランク: D（ギリギリ！）")
	}
}