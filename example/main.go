package main

import (
	"fmt"
	"example/math" 
	"util"
)

func main() {
    fmt.Println("Hello world! My lucky number is", math.Mul2(privateHelperFunc()))
    fmt.Println("Square:::", util.Square(60))
}
