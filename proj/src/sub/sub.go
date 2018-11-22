package sub

import . "fmt"

func init() {
	Println("sub is initial------")

}

func Add(a, b int) int {
	Print("this is add\n")
	return a + b
}
