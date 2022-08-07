package main

import "encoding/json"

func JsonString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
