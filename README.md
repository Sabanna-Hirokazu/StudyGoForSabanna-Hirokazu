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
Goの変数の宣言が今までにない感じで新しい。基本的に"var '変数名' '型'"で宣言できる。難しくはないが、なれる必要があると感じた。それに加え型推論がとてもしっかりとしているため、"変数名:=12"というやり方でもint型と定義される。"var 変数名 = 10"と"変数名 := 10"はどっちが多く使われているのだろう。([variableTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/variableTest.go))

### 3日目
本日はif文とswitch文に取り組んだ。if文の使い方が少しむずかしい。基本的にelseやelse ifを書くときはif文で閉じた中括弧の後ろに改行せずにする必要がある。switch文はcaseのあとに複数値を入力することが出来る。複数入力するとorの関係性と等しい。([ifTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/ifTest.go))

```
if x < 10 {
    println("ifのあとは改行してはいけない")
} else {
    println("elseは改行ぜずにいれます")
}
```

### 4日目
for文について学んだ。Goのfor文は少し特別な気がした。Goにはwhile文がないようなのでwhileの代わりにfor文を使う。for文に引数を与えないと無限ループになる。次にfor文の後ろに条件式を記載することで条件式を満たすうちはfor文が回るようになる。その他にもC言語などのように値の宣言から条件式や回ってきた際の処理をいれるということも可能。最後のfor文は0,1,2,3,4と表示され,jのfor文の中身はj=0で処理されるようにcontinue "Label name"にすることで"Label name"を記載したところから始めることが出来る。([forTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/forTest.go))

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
arrayとsliceについて学んだ。この2つはどちらも配列を表していて、arrayは初期化の際に定めたサイズより大きい配列にすることができない。それに比べsliceは可変長の配列であることがわかった。配列の中に配列を入れることで多次元配列も再現することができる。arrayで長さを指定するときに...にすることで自動的にサイズを指定することができる。値を追加したいときはsliceでのみappendを使用する。追加の関数が用意されているが値の消去の関数は存在しない。自分自身で作るしかない。([listTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/listTest.go))

```
// arrayを作る
array := [3]int{1,2,3} もしくは array := [...]int{1,2,3}
// sliceを作る
slice := []int{1,2,3}
// 値の追加(sliceにvalueを追加する)
slice := append(slice, value)
```

### 6日目
sliceに値を取り除くものが存在していなかった。そのため今回は値を取り出すための関数removeを定義してみた。goの...を使うのがとても便利であることを知った。今回は取り出す位置の手前をsliceとして抽出し、取り出す位置よりあとのsliceをappendすることでremoveの関数を実現した。([removeSlice.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/removeSlice.go))

### 7日目
今回はswitchについて学んだ。基本的な部分はifやelse ifといったものと同じ。"switch 条件式"というだけでなく、"switch 変数"という形にもすることができる。下記のようにすることができ、caseにおける値は一つでも複数でも可能。この場合"2, 3"というのは"if count == 2 || count == 3"というようなorの条件式を設定することが可能。([switchTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/switchTest.go))

```
count := 4
switch count {
    case 2, 3:
        println(23)
    case 4, 5:
        println(45)
    default:
        println(0)
    }
```

### 8日目
今回はWebのフレームワークginについて学んでいく。GoのWebのフレームワークで人気のようなので学んでみます。[gin](https://github.com/gin-gonic/gin)の通りにインストールなどをすすめる。gin.Default()のGETを指定するだけでルーティングすることが出来る。([exampleGin.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/exampleGin.go))

### 9日目
前回に引き続きGinについて学ぶ。WebのフレームワークといったらHtmlを表示したくなる。HTML renderingの説明にある文を用いて作成を行った。htmlを扱うときはimportを忘れずにする。基本的に読み込みたいhtmlファイルなどは先に読み込みをしておく必要がある。フォルダ内のファイルをすべてを使うときにLoadHTMLGlobが役立つ。＊などを使うことが出来るからだ。そのかわりに複数のものを指定することは出来ないので、全てではなく一部を使いたいときはr.LoadHTMLFilesを使うと良い。これは＊などですべてのファイルを指定することは出来ないが、複数のディレクトリのパスを指定することが出来る。必要に応じて使用する必要がある。Gin Debugを見ると読み込んだファイルを確認することが出来る。([sampleHtml.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/sampleHtml.go))


### 10日目
Ginについて学ぶ。Basic認証が使える。これはhttpなどの通信を簡単に認証方式を導入することが出来る。サンプルコードを一部変更しただけだが、データベースにあたる変数secretsがあるため、データベースの値をうまく整えることでデータベースとの連携も出来ると考える。([sampleBasicAuth.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/sampleBasicAuth.go))


### 11日目
Ginを学んでいく上でif文の新たな使い方を知った。シンプルステートメントというらしい。ifにはtrue、falseのように真偽値が必要だが、ifの関数内だけで変数を使いたいというときが来るかもしれない。値の宣言や代入などは真偽値の前に置き、セミコロンで閉じることで実行できるようになる。([ifTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/ifTest.go))


### 12日目
Goでmap関数を用いて辞書のようなものを作成した。Ginを学んでいるときにjsonのような形式の価をどのように行っているかの理解を深めるために作製した。辞書型のようは配列を作成し、キーを用いて価を参照すると戻り地に価が２つ受け渡されることに注意が必要！([dictionaryTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/dictionaryTest.go))


### 13日目
GinでBasic認証を用いるのは難しい。BasicAuth()に関するものがあまりなく、例外に対応することが出来なかった。BasicAuth()によって定められたアカウントのキーとは異なる価が入力をされたとする。しかしfunc(c *gin.Context)の中身を実行してくれない。そのため、特定のユーザーとパスワードが一致したときに関数は実行されるが、一致しないときには関数を実行しくれない。他にも実行方法があるのかを検討中。([sampleBasicAuth.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/sampleBasicAuth.go))

### 14日目
GinでQueryの処理をやってみた。今回はGETで値を取ってくるように作成した。curlの場合はバックスラッシュが無いと複数パラメータを投げれないことがわかった。これである程度のWebアプリケーションの実装が可能になりそう。([sampleQuery.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/sampleQuery.go))

```
// ブラウザなど
http://localhost:8080/get?id=10&page=1&name=hoge&message=fuga
//　ターミナルやコマンドプロンプト
curl localhost:8080/get?id=10\&page=1\&name=hoge\&message=fuga
```


### 15日目
前回と同じようにクエリをやってみる。今回はPOSTでチャレンジ！これでフォームに使われる処理をある程度理解できた。([sampleQuery.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/SampleGin/sampleQuery.go))

```
curl -X POST -F "name=hoge" -F "message=Fuga" localhost:8080/post?id=10\&page=1
```


### 16日目
ポインタについて学ぶ。C言語と同じように扱うことが出来る。変数に＆を付けることでその変数のアドレスを取得できる。アドレスを保存しておくためにはポインタを使う。変数の型の前に＊を付けるだけでその型のポインタになる。今回試したのはn1のアドレスを指しているのがn2という環境で行い、もちろんn2はn1を指し示しているため、n1の価が変更されればn2の価も変更されていることがわかった。([pointerTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/pointerTest.go))


### 17日目
今回は構造体を学ぶ。Goにも構造体が存在することがわかった。構造体を作成するためには方法が２つあり、new()を用いて初期化する方法と{}を用いて初期化する方法がある。newでは()の中身に構造体の名前を入れることで作成することが出来る。オブジェクト指向のコンストラクタに近い。{}の場合は価を代入した上で初期化することが出来るのでこちらのほうがよく使う印象。newはポインタで生成れることに対し、{}はないことに注意。構造体に関数を付与することが出来る。関数名の前に（引数　構造体名）で構造体の関数として定義することが出来る。値の計算などには問題ないが、宣言した構造体の値に代入する必要がある場合はポインタを用いる必要がある。構造体を入れこのように出来るが、ポインタなどの型に当てはめる必要がある。([structTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/structTest.go))


### 18日目
Goのパッケージについて学ぶ。オブジェクト指向のようにカプセル化をすることで、必要な操作のみを可能にさせ、中身を隠すことが出来る。使用する場合はパッケージの名前をフォルダにつけ、そのフォルダの階層内にファイルを追加していく。ここでフォルダの名前は最初に大文字を使うと使用することが出来なかった。Javaなども小文字が推奨されているっぽい。フォルダ内にあるGoファイルは"package foldername"にする必要がある。定義する関数は大文字始まりでないと認識することが出来ない。外部から使用する場合はフォルダ名をインポートすれば使用することが出来る。([sortTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/sortTest.go))


### 19日目
ソートがどうやらGoでは標準搭載ではないらしい。そのため、ソートしたいときのためにソートの関数を作成する。前回のパッケージを用いて作成する。今回はバブルソートを実装した。バブルソートはリストの先頭から進み、次の値と比較する。次の値より大きい場合は、その値と次の価を入れ替える。それを繰り返すことで最後に一番大きいもしくは小さい価がくる。これを配列のサイズの数だけ回す。([BubbleSort.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/mysort/BubbleSort.go))


### 20日目
ソートしたいときのためにソートの関数を作成する。前回と同様に作成する。今回は選択ソートを実装した。選択ソートはリストの先頭と後ろの値を比較する。後ろの値より大きい場合は、その値と後ろの値を入れ替える。このとき入れ替える後ろの値は、後ろの値の中で最小値もしくは最大値を選ぶ。それを繰り返すことで最初に一番大きいもしくは小さい値がくる。これを配列のサイズの数だけ回す。([SelectionSort.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/mysort/SelectionSort.go))


### 21日目
挿入ソートを実装した。挿入ソートは、先頭の値はソート済みと考え、残りの次の値をどこに入れるかをソート済みの後ろの値から比較して考える。これを繰り返すことでソートすることができる。([InsertionSort.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/mysort/InsertionSort.go))


### 22日目
マージソートを実装した。マージソートは半分で分割を行い、その分割を行ったものをサイズの順番にソートしながら結合するアルゴリズム。[MergeSort.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/mysort/MergeSort.go))

### 23日目
クイックソートを実装した。クイックソートはピボットと呼ばれる値を取り、ピボットより小さいかもしくはピボット以上かでグループ分けを行う。ピボットの値によって計算量が大きく変わってくる。今回は先頭の値もしくは次の値を取るクイックソートで実装した。グループを分けていく過程で、ピボットが最小値の場合、グループ分けができないので、次の値が大きい場合は、次の値をとるようにした。[QuickSort.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/mysort/QuickSort.go))


### 24日目
今回はInterfaceを学んだ。interface自体はJavaのinterfaceとほとんど考え方は同じだが、実装するimplementsの宣言がないためわかりにくい。interfaceには共通して必要なものを宣言だけしておく。その後実装する関数でその内容を記載する必要がある。Goの場合、interfaceの型に宣言した変数にキャストしていれるようなイメージで作成することができる。interfaceで宣言したもの以外を宣言する、もしくは宣言したものを実装していない場合エラーが出るので注意が必要だ。[interfaceTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/interfaceTest.go))


### 25日目
本日はGoの並列処理について学んでいく。Goでは並列処理がとてもやりやすいときいていたがとても奥が深いと感じた。Goの並列処理にはchannelとWaitGroupを使うパターンが存在した。今回は並列にするのと並列にしないとではどれくらい時間が変わるのかを計測した。forループで計測した。pythonのようにpassみたいなものがなかったため書き方が少し雑にはなったが検証することが出来た。並列を行わず関数を２回呼び出した結果は0.8306973秒かかった。しかし並列で関数を実行することで0.435357372になった。約半ほど時間が減ったことがわかる。多次元配列などのソートなどにおいてかなり役立つのではないかと思った。[parallelProcesingTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/parallelProcesingTest.go))


### 26日目
今回はGoの無名関数について学んでいく。Goの並列処理を学んでいたとき、よくわからない関数が登場したため勉強をする。前回行われていた関数の表記はs4のやり方と等しい。基本はmainの関数外で関数を宣言して行うが、無名関数はその関数内で宣言することが出来る。通常の関数は関数の中に宣言することが出来ない。いろいろな方法があるのではないかと考える。[functionTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/functionTest.go))


### 27日目
今回はGoのコールバック関数とクロージャについて学んでいく。コールバック関数は引数に関数をいれること。クロージャは関数内で宣言された変数を参照することが出来る。シングルトンのように外部からの値の変更を防ぐことが出来ると思われる。関数を呼び出す際に使う変数を通常使用するときはグローバル変数などを用いるが、関数内のみで変数を扱うことが出来るので、グローバル変数で使うより、値などが変わってしまう心配が減るみたい。[functionTest.go](https://github.com/Sabanna-Hirokazu/StudyGoForSabanna-Hirokazu/blob/master/functionTest.go))