# json_filter


```go
package main
import json-filter

func TestFilter() {
	var env interface{}
	err := json.Unmarshal([]byte(`{"email":"rsj217@gmail.com","money":100.5,"a":{"b":"bbbb","c":{"d":"ddddd"}},"e":"eeee"}`), &env)
	re, err := Filter(env, "a.c.h,email")
	//json转结构体
	result, err := json.Marshal(re)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(string(result))
}
```

