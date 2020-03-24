package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"io"
	"io/ioutil"
	_ "lang/calc"
	"lang/myLog"
	"os"
	"time"
)

const (
	len int = 124
)

func fileTobufio() {
	fh, err := os.Open("./calc/calc.go")
	if err != nil {
		fmt.Println("file open failed")
		return
	}
	defer fh.Close()
	reader := bufio.NewReader(fh)
	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("读取到末尾")
			return
		}

		if err != nil {
			fmt.Println("read failed")
			return
		}

		fmt.Println(str)
	}

}

func readByIoUtil(str string) {
	tmp, err := ioutil.ReadFile(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(tmp))

}

func writeDemo1() {
	fh, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	for i := 0; i <= 4; i++ {

		n, err := fh.Write([]byte("hello 岳伟超\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n)
	}

}

func writeDemo2() {
	fh, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()

	for i := 0; i < 5; i++ {
		n, err := fh.WriteString("hello,葛海琴\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n)
	}

}

func writeDemo3() {
	//fh, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer fh.Close()
	err := ioutil.WriteFile("./test", []byte("hahhaa,kkksldk\nhhjajajhahs\n"), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func test(str string, a ...interface{}) {
	msg := fmt.Sprintf(str, a)
	fmt.Println(msg)
}
func main() {
	/*	pc, file, line, ok := runtime.Caller(0)
		if ok{
			fmt.Println(strings.Split(runtime.FuncForPC(pc).Name(),".")[0],path.Base(file),line)
		}*/
	log := myLog.NewFileLogger("./path/", "log", "debug", 10*1024*1024)
	//str := "yueyue%d"
	//test(str,4)
	for i:=0;i<100;i++ {
		log.Debug("这是一个debug测试")
		log.Trace("这是一个trace测试")
		name := "yue"
		age := 10
		log.Info("这是一个info测试 name[%s] age[%d]", name, age)
		log.Warning("这是一个warning测试")
		log.Error("这是一个error测试")
		log.Fatal("这是一个fatal测试")
		time.Sleep(time.Duration(1 * time.Second))
	}
	log.FClose()
	//writeDemo2()
	//writeDemo3()
	//writeDemo1()
	//fileTobufio()
	//readByIoUtil("./calc/calc.go")
	//fmt.Println(calc.Jc(3))
	/*fh, err := os.Open("./calc/calc.go")
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
	}*/

}
