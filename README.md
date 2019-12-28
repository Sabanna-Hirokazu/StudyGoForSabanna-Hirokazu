<!-- ![] (http) -->

# StudyGoForSabanna-Hirokazu
<b>Study Go</b>
新たにGoという言語を勉強するためのプロジェクト

## 目的
最近はかなり聞くようになり、ベンチャー企業などでは新しいGoを取り入れようという動きがあると就活を通して知り、学ぶ

## 実行環境
- OS: Zorin 15 core (Ubuntu 18.04 派生)
- Go: gol.13.5 linux/amd63
- Browser: Chrome

## 日々の成果

### 1日目
GoはGoogleの人が作製した言語であり、静的な型付けが行われるらしい。並列処理も簡単に出来るようだ。本日はGo言語のインストールから実際にコンパイルまでを取り組む。実行時には下記のようにやる必要がある。

```
//一度コンパイルをしてファイルを作成する
go build "FileName"
./ "Filename"

//まとめて実行する
go run "FileName"
```

### 2日目
Goの変数の宣言が今までにない感じで新しい。基本的に"var '変数名' '型'"で宣言できる。難しくはないが、なれる必要があると感じた。それに加え型推論がとてもしっかりとしているため、"変数名:=12"というやり方でもint型と定義される。"var 変数名 = 10"と"変数名 := 10"はどっちが多く使われているのだろう。

### 3日目
本日はif文とswitch文に取り組んだ。if文の使い方が少しむずかしい。基本的にelseやelse ifを書くときはif文で閉じた中括弧の後ろに改行せずにする必要がある。switch文はcaseのあとに複数値を入力することが出来る。複数入力するとorの関係性と等しい。

```
if x < 10 {
    println("ifのあとは改行してはいけない")
} else {
    println("elseは改行ぜずにいれます")
}
```

### 4日目
for文について学んだ。Goのfor文は少し特別な気がした。Goにはwhile文がないようなのでwhileの代わりにfor文を使う。for文に引数を与えないと無限ループになる。次にfor文の後ろに条件式を記載することで条件式を満たすうちはfor文が回るようになる。その他にもC言語などのように値の宣言から条件式や回ってきた際の処理をいれるということも可能。最後のfor文は0,1,2,3,4と表示され,jのfor文の中身はj=0で処理されるようにcontinue "Label name"にすることで"Label name"を記載したところから始めることが出来る。

```
for{
    println("無限ループ")
}
for count < 5 {
    println("countが5になるまでまわす")
}
for i:= 0; i<5; i++{
    println("c言語のような感じ")
}
Label:
for i:= 0; i<5; i++{
    for j:= 0; j<5; j++{
        println(i+j)
        continue Label
    }
}
```

### 5日目
arrayとsliceについて学んだ。この2つはどちらも配列を表していて、arrayは初期化の際に定めたサイズより大きい配列にすることができない。それに比べsliceは可変長の配列であることがわかった。配列の中に配列を入れることで多次元配列も再現することができる。arrayで長さを指定するときに...にすることで自動的にサイズを指定することができる。値を追加したいときはsliceでのみappendを使用する。追加の関数が用意されているが値の消去の関数は存在しない。自分自身で作るしかない。

```
// arrayを作る
array := [3]int{1,2,3} もしくは array := [...]int{1,2,3}
// sliceを作る
slice := []int{1,2,3}
// 値の追加(sliceにvalueを追加する)
slice := append(slice, value)
```

### 6日目
sliceに値を取り除くものが存在していなかった。そのため今回は値を取り出すための関数removeを定義してみた。goの...を使うのがとても便利であることを知った。今回は取り出す位置の手前をsliceとして抽出し、取り出す位置よりあとのsliceをappendすることでremoveの関数を実現した。([removeSlice.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/removeSlice.go)に記載)