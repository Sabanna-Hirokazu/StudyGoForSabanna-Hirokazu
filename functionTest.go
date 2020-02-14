package main

func sample(tmp string) string {
	return "'" + tmp + "'"
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
}
