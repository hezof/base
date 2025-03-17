package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Score  float32 `json:"score"`
	Female bool    `json:"female"`
	Secret []byte  `json:"secret"`
}

type Class struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Student struct {
	*Person
	Class  *Class    `json:"class"`
	Parent []*Person `json:"parent"`
}

var mdata = map[string]interface{}{
	"name":   "zenlili",
	"age":    99,
	"score":  113.5,
	"female": true,
	"secret": "测试数据",
	"class": map[string]interface{}{
		"name": "三班",
		"code": "sanban",
	},
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
	var stu *Student
	err := MapStruct(mdata, &stu, "json")
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(stu)
	fmt.Printf("Data: %s\n", data)
}
