/*
Uy vazifa
1
Goroutine lar  va channel lar yordamida bir vaqtning o'zida berilgan sonning faktorialini hisoblaydigan Go dasturini yozing.
*/

package main

import "fmt"

func FindFactorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func main() {
	num := 5

	ch := make(chan int)

	go FindFactorial(num, ch)

	result := <-ch

	fmt.Printf("%d ning faktoriali: %d\n", num, result)
}
