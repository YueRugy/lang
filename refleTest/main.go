package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	. "strings"
)

/*type person struct {
	Name string `json:"name"`

	Age int `json:"age"`
}*/
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database int    `ini:"database"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func main() {
	var mc Config

	err := loadIni("./conf.ini", &mc)
	fmt.Println(err)
}

func loadIni(fileName string, data interface{}) error {
	//0 参数的校验
	//01 data 必须是指针类型
	t := reflect.TypeOf(data)
	//fmt.Println(t.Elem().Kind())
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return errors.New("not struct ptr type")
	}
	//2读取文件
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	var structName string
	//2.1 一行一行读数据
	strArray := Split(string(b), "\n")
	for idx, line := range strArray {
		line = TrimSpace(line)
		//2.2注释跳过
		if HasPrefix(line, ";") || HasPrefix(line, "#") ||
			len(line) == 0 {
			continue
		}
		//如果是[开头表示节(section)
		if HasPrefix(line, "[") {
			//校验是否是合法的section
			//左右中括号不匹配
			if !HasSuffix(line, "]") {
				return fmt.Errorf("line %d syntax error", idx+1)
			}

			sectionName := TrimSpace(line[1 : len(line)-1])
			//中括号内没有内容
			if len(sectionName) == 0 {
				return fmt.Errorf("line %d syntax error", idx+1)
			}

			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					structName = field.Name
					fmt.Println(structName)
					break
				}
			}
		} else {
			if Index(line, "=") == -1 || HasSuffix(line,
				"=") || HasPrefix(line, "=") {
				return fmt.Errorf("line %d syntax error", idx+1)
			}

			v := reflect.ValueOf(data)
			sv := v.Elem().FieldByName(structName) //拿到嵌套结构体值信息
			st := sv.Type()                        //拿到嵌套结构体类型信息
			if st.Kind() != reflect.Struct {
				return fmt.Errorf("data 中的%d字段应该是一个结构体",
					idx+1)

			}

			index := Index(line, "=")
			k := line[:index]
			//value := line[index+1:]
			var fieldType reflect.StructField
			for i := 0; i < st.NumField(); i++ {
				//field := sv.Field(i)
				field := st.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == k {
					fileName = field.Namegit
					break
				}

			}
			fmt.Println(fileName,fieldType.Type.Kind())

		}
	}
	//fmt.Println(strArray)

	return nil

}
