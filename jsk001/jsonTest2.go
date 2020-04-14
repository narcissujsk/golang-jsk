package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

//map转json
func mapToJSON(tempMap *map[string]interface{}) string {

	data, err := json.Marshal(tempMap)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func Testjson(t *testing.T) {

	tempMap := make(map[string]interface{})

	tempMap["username"] = "itbsl"
	tempMap["age"] = 18
	tempMap["sex"] = "male"

	str := mapToJSON(&tempMap)

	fmt.Println(str)

}
