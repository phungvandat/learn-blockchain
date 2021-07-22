package main

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(val interface{}) {
	b, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
