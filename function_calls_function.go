package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
	f3(a)
}

func f2() {
	print(a)
}
func f3(a string) {
	print(a)
}
