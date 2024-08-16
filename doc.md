# 雑メモ

## 用語

#### レシーバ: メソッドが属する構造体を示す（`(p *Person)`の`p`のこと）

```go
type Person {
    name: string
    age: int
}

func (p *Person) Greeting() {
  fmt.Println("Hello!")
}
```

#### ポインタ: その変数の値が格納されている場所（アドレス）を直接参照するもの

`ポインタを使用する場合`: 渡された構造体の値に直接影響を与える
`ポインタを使用しない場合`: 構造体のコピーを作ってからメソッドに構造体を渡す挙動になる

```go
// ポインタを使用する場合
func (p *Person) addAge() {
    p.age += 1
}

func main() {
    alice := Person{name: "Alice", age: 18}
    fmt.Println(alice) // {"Alice", 18}

    alice.addAge()

    // 構造体の値のアドレスを直接参照し、ageを変更するため、元の構造体に影響を与える
    fmt.Println(alice) // {"Alice", 19}
}
```

```go
// ポインタを使用しない場合
// コピーして同じ値のPerson型を作ってからメソッドに構造体を渡す
func (p Person) addAge() {
    p.age += 1
}

func main() {
    alice: Person{name: "Alice", age: 18}
    fmt.Println(alice) // {"Alice", 18}

    // コピーした構造体に対してaddAge()を実行
    alice.addAge()

    // コピーした構造体に対してaddAge()を実行したため、元の構造体には影響がない
    fmt.Println(alice) // {"Alice", 18}
}
```

#### パッケージ：　同じディレクトリの中でまとめられた、変数や定数、関数定義の集合のこと

#### モジュール: パッケージの集まりのこと

#### ハンドラ: HTTP リクエストを受け取って、それに対する HTTP レスポンスの内容をコネクションに書き込む関数のこと

```go
// 必ず以下の形にする必要がある
// 引数: http.ResponseWriter型とhttp.Request型をとる
// 戻り値: なし
func (w http.ResponseWriter, req *http.Request) {}
```
