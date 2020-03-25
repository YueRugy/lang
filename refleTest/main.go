package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
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
	//fmt.Println(err)
	if err != nil {
		print("haha")
	}
	fmt.Println(mc)
}

func loadIni(fileName string, data interface{}) error {
	//判断data 是否是结构体的指针类型
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Ptr || dataType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("syntax type data")
	}

	//读文件
	fileByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("open ini file failed")
	}
	//转换成字符串
	fileContent := string(fileByte)
	contentArray := strings.Split(fileContent, "\n")
	var sectionName string
	var structName string
	//一行行解析字符串
	for idx, line := range contentArray {
		//注释空行跳过
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		//[表示一个节点section
		if strings.HasPrefix(line, "[") {
			//校验合法性
			//获取节点名
			sectionName = strings.TrimSpace(line[1 : len(line)-1])
			if !strings.HasSuffix(line, "]") ||
				len(sectionName) == 0 {
				return fmt.Errorf("%d行syntax error section", idx+1)
			}

			for i := 0; i < dataType.Elem().NumField(); i++ {
				field := dataType.Elem().Field(i)
				tag := field.Tag.Get("ini")
				if tag == sectionName {
					structName = field.Name
					break
				}
			}

		} else {

			//在dataType获取与sectionName相同的tag
			sectionValue := reflect.ValueOf(data).Elem().FieldByName(structName)
			//fmt.Println(sectionValue.Kind())
			sectionType := sectionValue.Type()
			//fmt.Println(sectionType.Kind())
			//是否是结构体
			if sectionType.Kind() != reflect.Struct {
				return fmt.Errorf("Config 中的变量 %v 不是结构体", sectionType.Name())
			}
			//获取=k ,v并校验
			if strings.Index(line, "=") == -1 || strings.
				HasSuffix(line, "=") || strings.
				HasPrefix(line, "=") {
				return fmt.Errorf("= syntax error %d", idx+1)
			}
			KVArray := strings.Split(line, "=")
			k := KVArray[0]
			v := KVArray[1]
			var fieldName string
			var field reflect.StructField
			//fmt.Println(k,v)
			//赋值
			for i := 0; i < sectionType.NumField(); i++ {
				field = sectionType.Field(i)
				if field.Tag.Get("ini") == k {
					fieldName = field.Name
					break
				}

			}

			if v == "jj" {

			}
			fmt.Println(fieldName, field.Type.Kind())
			sectionValueField := sectionValue.FieldByName(fieldName)
			//	fmt.Println(field.Name,field.Type.Kind())
			switch field.Type.Kind() {
			case reflect.String:
				sectionValueField.SetString(v)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, _ := strconv.ParseInt(v, 10, 64)
				sectionValueField.SetInt(value)
			}

		}

		//fmt.Println(idx, line)

	}

	return nil
}
