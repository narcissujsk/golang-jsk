package jsk001

import (
	"encoding/json"
	"fmt"
	"testing"
)

//map转json
func mapToJSON(tempMap *map[string]interface{}) string {

	data, err := json.Marshal(tempMap)

	if err != nil {

	}

	return string(data)
}
func TestMapToJson(t *testing.T) {
	tempMap := make(map[string]interface{})

	tempMap["username"] = "itbsl"
	tempMap["age"] = 18
	tempMap["sex"] = "male"

	str := mapToJSON(&tempMap)

	fmt.Println(str)
}
