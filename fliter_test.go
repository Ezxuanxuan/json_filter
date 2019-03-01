package json_filter

import (
	"encoding/json"
	"fmt"
	"testing"
)

//func TestFilter(t *testing.T) {
//	jsons := Test{
//		email: "rsj217@gmail.com",
//		money: 1000,
//		d: "dddddd"
//	}
//
//	Filter(jsons, "email,a.b.c,money")
//}

func TestFilter(t *testing.T) {
	var env interface{}
	err := json.Unmarshal([]byte(`{"email":"rsj217@gmail.com","money":100.5,"a":{"b":"bbbb","c":{"d":"ddddd"}},"e":"eeee"}`), &env)
	re, err := Filter(env, "a.c.h,email")

	//返回结构转为json
	result, err := json.Marshal(re)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(string(result))
}

type Test struct {
	email string
	money int32
	a     interface{}
	d     string
}
