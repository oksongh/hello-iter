## 何をした？
Goのイテレータを理解するためにサンプルコードを写経した。

## 何を理解した？
for rangeに見えるけど、実態としてはループしてない、けど気持ち的にはループしている構文(?)、 range over function が追加された。
range over functionの見た目はnumbersから取り出してvに代入し、`fmt.Print`で表示しているように見える。
使用者側からはこの認識で問題ないが、実際にループしているのは`numbers`の内部。

また、`numbers`の型、`func(yield func(int) bool)`はよく使われるので、iter.Seq[E]とジェネリクスで定義されていて、シーケンスと呼ばれている。

シーケンスはその名の通り、リストやファイルの中身のような`一連の値`を取り出せて、それをrange over forで使う、という流れになる。

```go
numbers := func(yield func(int) bool) {
    // 本物のループ
    for i := range 10 {
        yield(i)
    }
}

// 実態としてはループじゃない。クロージャを渡して実行している。
for v range numbers {
    fmt.Printf("hey %v\n", v)
}
```

## range over functionの実態
これは調べてないので想像だが、range over functionは糖衣構文で、実態としてはfor rangeの部分が関数にラップされてnumbers(func()...)のようになっていそう。

もちろんbreakやcontinue、returnといった制御構文も使える。
脱糖するときに`return false`なんかに置換されていそうな雰囲気。

## 結論
* 面白いのは木構造やグラフに対してシーケンスを定義してやると単一のforで全探索ができたりするところ。
* はじめはRustのイテレータのアプローチのほうが良いのでは？と思っていた。
が、使ってみると思ったより機能していそうな感触がある。

* forが range over forなのか、実態があるのかを区別しながら読むべきかはわからなかった。
区別しないといけないなら可読性が微妙だが、どうだろう？
