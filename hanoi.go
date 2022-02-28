package main

func main(){
	Towers(4, "Pike 1", "Pike 2", "Pike 3")
}
func print_move(fr string, to string) {
	println("move from ", fr, " to ", to)
}
func Towers(n int, fr string, to string, spare string) {
	if n == 1 {
		print_move(fr, to)
	} else {
		Towers(n-1, fr, spare, to)
		Towers(1, fr, to, spare)
		Towers(n-1, spare, to, fr)
	}

}