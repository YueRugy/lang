package main

import (
	"fmt"
	_ "fmt"
	_ "lang/calc"
	"os"
)

const (
	len int = 124
)

func main() {

	//fmt.Println(calc.Jc(3))
	fh, err := os.Open("./calc/calc.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	defer fh.Close()

	var buffer [128]byte
	for {
		n, err := fh.Read(buffer[:])
		if err != nil {
			fmt.Println("file read failed")
			return
		}
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(buffer[:n]))
		if n < 128 {
			return
		}
	}

}
