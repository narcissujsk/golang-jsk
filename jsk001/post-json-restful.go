package main

import (
	"fmt"
	"net/http"
	_ "strings"
)

import (
	"bytes"
	"io/ioutil"
)

func postjson(url string,body string) string{
	var jsonStr = []byte(body)
	fmt.Println("jsonStr", string(jsonStr))
	fmt.Println("new_str", bytes.NewBuffer(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	re, _ := ioutil.ReadAll(resp.Body)
	return string(re)

}

func main() {
	url := "http://192.168.152.1:8080/metrics/cps"

	fmt.Println("url:>", url)

	usrId := "login"
	pwd := "pwd12344"

	//json序列化
	post := "{\"UserId\":\"" + usrId +
		"\",\"Password\":\"" + pwd +
		"\"}"

	fmt.Println(url, "post", post)
    dd:=postjson(url,post)
	fmt.Println("response Body:", dd)
}
