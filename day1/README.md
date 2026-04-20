# Day 01 — Hello World CLI

## 学んだこと

- Go の基本構造（package / import / func main）
- `go mod init` で Go Modules プロジェクトを作成
- `go run` と `go build` の違い
- `os.Args` でコマンドライン引数を取得
- `os.Getenv` で環境変数を取得
- `fmt.Println` と `fmt.Printf` の違い
- `:=`（宣言+代入）と `=`（再代入）の違い

## 使い方

```bash
go run main.go                              # Hello, World!
go run main.go Taro                         # Hello, Taro!
GREETING=こんにちは go run main.go 太郎       # こんにちは, 太郎!
```

## ハマった点
main関数のエントリポイントがあるファイルのpackageはmainにしないと動かない。

## 次回改善したいこと

