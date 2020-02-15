package main

var k int = 4

func sample(tmp string) string {
	return "'" + tmp + "'"
}

// CallBack  第二引数には引数が文字列、戻り地が文字列の関数が入る
func sample2(tmp string, f func(string) string) string {
	return f("sample2 - " + tmp)
}

// クロージャ
func sample3() {
	var i int = 4
	f := func() {
		i--
		println(i)
	}
	for i > 1 {
		f()
	}
}

func main() {
	// 通常の関数
	s1 := sample("Hoge")
	println(s1)
	f1 := sample
	s2 := f1("Fuga")
	println(s2)

	// 無名関数
	f2 := func(tmp string) string { return "''" + tmp + "''" }
	s3 := f2("Hello")
	println(s3)
	s4 := func(tmp string) string { return "+" + tmp + "+" }("OK")
	println(s4)

	// コールバック関数
	s5 := sample2("World", sample)
	println(s5)

	// クロージャ
	// pattern 1
	sample3()
	// pattern 2
	for k > 1 {
		k--
		println(k)
	}
}
