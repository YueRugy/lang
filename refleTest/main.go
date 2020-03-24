package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`

	Age int `json:"age"`
}

func main() {
	var p person
	str := `{"name":"yue","age":18}`
	_ = json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
