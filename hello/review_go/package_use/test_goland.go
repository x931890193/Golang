package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Printf("os.Args's type is %T, length is %d\n", os.Args, len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
	//os.Chdir("./")
}

func test() {
	//pass
}
