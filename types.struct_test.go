package core

import (
	"fmt"
	"testing"
)

type Person struct {
	Name   string  `json:"name,omitempty"`
	Age    int     `json:"age,omitempty"`
	Score  float32 `json:"score,omitempty"`
	Female bool    `json:"female,omitempty"`
	Secret []byte  `json:"secret,omitempty"`
}

type Class struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type Student struct {
	*Person
	Class  *Class    `json:"class,omitempty"`
	Parent []*Person `json:"parent,omitempty"`
}

var mdata = map[string]interface{}{
	"name":   "zenlili",
	"age":    99,
	"score":  113.5,
	"female": true,
	"secret": "测试数据",
	"class":  nil,
	"parent": []interface{}{ // 注意: 泛型slice是[]interface{}
		map[string]interface{}{
			"name": "baba",
		},
		map[string]interface{}{
			"name": "mama",
		},
	},
}

func TestMapStruct(t *testing.T) {

	fmt.Println(ToJson(mdata))
	var stu *Student
	err := MapStruct(mdata, &stu, "json")
	if err != nil {
		panic(err)
	}
	fmt.Println(ToJson(stu))
	bs, _ := UnBase64("5rWL6K+V5pWw5o2u")
	fmt.Println(string(bs))
}
