package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// スキャナーの準備
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) == 4 {
			num1, err1 := strconv.ParseFloat(os.Args[1], 64)
			num2, err2 := strconv.ParseFloat(os.Args[3], 64)
			operator := os.Args[2]

		if err1 != nil || err2 != nil {
				fmt.Println("エラー: 引数に正しい数値を指定してください")
				os.Exit(1)
		}

		result, calcErr := calculate(num1, num2, operator)
		if calcErr != nil {
			fmt.Println("エラー：", calcErr)
			os.Exit(1)
		}

		fmt.Printf("%g\n", result)
		return
	}

	fmt.Println("==== 簡易電卓 ====")
	fmt.Println("終了するには'q'を入力")
	fmt.Println()

	// forの無限ループ goにはwhileが無い
	for {

		fmt.Print("1つ目の数値 (q で終了): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if input == "q" || input == "quit" {
			fmt.Println("電卓は終了します")
			break
		}

		// 文字列をfloat64に変換
		num1, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("エラー：数値を入力してください")
			continue
		}

		// 演算子を入力
		fmt.Print("演算子 (+, -, *, /): ")
		scanner.Scan()
		operator := strings.TrimSpace(scanner.Text())

		// 2つ目の数値を入力
		fmt.Print("2つ目の数値: ")
		scanner.Scan()
		num2, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("エラー：数値を入力してください")
			continue
		}

		// 計算実行
		result, calcErr := calculate(num1, num2, operator)
		if calcErr != nil {
			fmt.Println("エラー：", calcErr)
			continue
		}

		// %gは不要な末尾ゼロを削除
		fmt.Printf("結果: %g %s %g = %g\n", num1, operator, num2, result)
	}
}

// 計算ロジック
// 慣習としてエラーの戻り値は最後におく
// 呼び出し側でエラーチェックする
func calculate(num1 float64, num2 float64, op string) (float64, error) {
		fmt.Printf("計算式の要素　num1：%f, op:%s, num2:%f\n", num1, op, num2)
		switch op {
		case "+":
				return num1 + num2, nil
		case "-":
				return num1 - num2, nil
		case "*":
				return num1 * num2, nil
		case "/":
				if num2 == 0 {
					return 0, fmt.Errorf("０で割ることはできません")
				}
				return num1 / num2, nil
		default:
			  return 0, fmt.Errorf("未対応の演算子です: %s", op)
		}
}