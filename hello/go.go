package main //

import "fmt"
import "time"

func main() {
	sum := 0
	st := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		//		fmt.Println(sum)
		sum += i

	}
	fmt.Println(sum)
	ed := time.Now().UnixNano()

	fmt.Println(ed - st)
}
