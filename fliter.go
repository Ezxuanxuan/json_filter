package json_filter

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

//一个proxy的json过滤嘴
/*
   env： 被解析的json字符串
   fields：需要获取的列名，用"，"隔开，多层需要用"."表示
   用例： env:`{"email":"rsj217@gmail.com","":100.5,"a":{"b":"bbbb","c":{"d":"ddddd"}},"e":"eeee"}`,  fields:"email,a.c.d,money"
*/
func Filter(env interface{}, fields string) (interface{}, error) {
	//var env interface{}
	////json转为结构体
	//if err := json.Unmarshal([]byte(jsons), &env); err != nil {
	//	return "", err
	//}
	// for the love of Gopher DO NOT DO THIS
	// 切割所有需要的字段名
	fies := strings.Split(fields, ",")
	//fmt.Println(env.(map[string]interface{})[fields].(string))

	//取出相应字段对应值
	re := make(map[string]interface{})
	var errs error
	for _, fie := range fies {
		re[fie], errs = get(env, fie)
		if errs != nil {
			return nil, errs
		}
	}
	return re, nil
}

/*
  { "a":"sss", "b":"sss", "c":{"d":"sss", "e":"sss"}, "e":"sss"}
*/

//结构体解析，递归取值
func get(env interface{}, fie string) (interface{}, error) {

	//不含有"."，直接取出对应的interface
	if !strings.Contains(fie, ".") {
		env1 := env.(map[string]interface{})
		_, ok := env1[fie]
		fmt.Println(ok)
		if !ok {
			return nil, errors.New("不存在字段:" + fie)
		}
		re := env.(map[string]interface{})[fie].(interface{})
		return re, nil
	} else {
		first := strings.Index(fie, ".") //判断第一个"."出现的位置
		//用于存取字段名，（上一层结构体）
		filedName := fie[0:first]
		//根据字段名获取结构体
		newEnv := env.(map[string]interface{})[filedName].(interface{})
		newFie := fie[first+1:]
		//递归
		return get(newEnv, newFie)
	}
}
