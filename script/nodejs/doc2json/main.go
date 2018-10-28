package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

func main() {
	b, err := ioutil.ReadFile("./resume.json")
	if err != nil {
		fmt.Printf("ReadFile err:%s\n", err.Error())
		return
	}

	var v interface{}
	err = json.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("json.Unmarshal err:%s\n", err.Error())
		return
	}

	v = decodeJSON(v)
	if err != nil {
		fmt.Printf("decodeJSON err:%s\n", err.Error())
		return
	}

	b, err = json.Marshal(v)
	if err != nil {
		fmt.Printf("json.Marshal err:%s\n", err.Error())
		return
	}

	ioutil.WriteFile("new.json", b, 0666)
}

func decodeJSON(obj interface{}) interface{} {
	switch obj.(type) {
	case map[string]interface{}:
		for k, v := range obj.(map[string]interface{}) {
			obj.(map[string]interface{})[k] = decodeJSON(v)
		}
	case []interface{}:
		for i, v := range obj.([]interface{}) {
			obj.([]interface{})[i] = decodeJSON(v)
		}
	case string:
		obj, _ = url.QueryUnescape(obj.(string))
		return obj
	}

	return obj
}
