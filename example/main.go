package main

import (
	"fmt"

	"github.com/KimoTek/Golang/example/mathutils"
	"github.com/KimoTek/Golang/util"
)

func main() {
	fmt.Println("Hello world! My lucky number is", mathutils.Mul2(privateHelperFunc()))
	fmt.Println("Square:::", util.Square(60))
}
