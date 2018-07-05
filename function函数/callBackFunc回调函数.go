package main

import "fmt"

/* 回调函数 :

 */
func main() {
	a1 := dt(3, 1, jiafa)
	a2 := dt(3, 1, jianfa)
	fmt.Printf("相加a1=%d,相减a2=%d\n", a1, a2)
}

type hs func(int, int) int

func dt(a, b int, js hs) (jg int) {
	jg = js(a, b)
	return
}
func jiafa(a, b int) int {
	return a + b
}
func jianfa(a, b int) int {
	return a - b
}
