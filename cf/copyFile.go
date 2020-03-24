package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	copyFile("./test.txt","test.txt")
 }
func copyFile(path, copyName string) {
	fh, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	copyName = "./path/" + copyName
	fh1, err1 := os.OpenFile(copyName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer fh.Close()
	reader := bufio.NewReader(fh)
	var tmp [128]byte
	for {
		buffer := tmp[:]
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		//wirte
		_, err = fh1.Write(buffer[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
		if n < 128 {
			return
		}
	}
}
