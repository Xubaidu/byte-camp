package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func QueryWithCaiyun(word string) []string {

	// construct request
	client := &http.Client{}
	data := map[string]interface{}{
		"trans_type": "en2zh",
		"source":     word,
	}
	buffer, err := json.Marshal(data) // serialize data to []byte
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(buffer) // pass bytes to reader

	// request
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	// response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body) // read bytes from reader
	if err != nil {
		log.Fatal(err)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(bodyText, &body) // deserialize bytes to map
	if err != nil {
		log.Fatal(err)
	}

	// parse response
	var ans []string
	explanations := body["dictionary"].(map[string]interface{})["explanations"].([]interface{})
	for _, v := range explanations {
		ans = append(ans, v.(string))
	}
	return ans
}

func QueryWithVolcano(word string) []string {

	// construct request
	client := &http.Client{}
	data := map[string]interface{}{
		"text":     word,
		"language": "en",
	}
	buffer, _ := json.Marshal(data)
	reader := bytes.NewReader(buffer)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDc7EuIJSWc-Ip7Tlqe8&_signature=_02B4Z6wo00001lWraGAAAIDBwX0zbJ9zEJZVq2zAAPcf58", reader)
	if err != nil {
		log.Fatal(err)
	}

	// request
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "_ga=GA1.2.425760739.1649834229; digest=foRRAkXrenF44o7aWYwZcPUfGPoUxsEcktmNB4EVwo8=; referrer_title=-%20%E6%96%87%E7%AB%A0%20-%20%E5%BC%80%E5%8F%91%E8%80%85%E7%A4%BE%E5%8C%BA; x-jupiter-uuid=16521869265494832; i18next=zh-CN")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=good")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

	// response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(bodyText, &body)

	// parse response
	var ans []string
	pos_list := body["words"].([]interface{})[0].(map[string]interface{})["pos_list"].([]interface{})
	for _, i := range pos_list {
		explanations := i.(map[string]interface{})["explanations"].([]interface{})
		for _, j := range explanations {
			ans = append(ans, j.(map[string]interface{})["text"].(string))
		}
	}
	return ans
}

func QueryWithEngine(word, engine string) []string {
	switch engine {
	case "caiyun":
		return QueryWithCaiyun(word)
	case "volcano":
		return QueryWithVolcano(word)
	default:
		return nil
	}
}

func QueryParallelly(word string, engines []string) []string {
	var ans []string
	var wg sync.WaitGroup
	start := time.Now()
	for _, engine := range engines {
		wg.Add(1)
		go func(engine string) {
			defer wg.Done()
			ans = append(ans, QueryWithEngine(word, engine)...)
		}(engine)
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println("Query parallelly, elapsed time:", end)
	return ans
}

func QuerySerially(word string, engines []string) []string {
	var ans []string
	start := time.Now()
	for _, engine := range engines {
		temp := QueryWithEngine(word, engine)
		ans = append(ans, temp...)
	}
	end := time.Since(start)
	fmt.Println("Query serially, elapsed time:", end)
	return ans
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	/*
		word := "good"
		engines := []string{"caiyun", "volcano"}
		ans := QueryParallelly(word, engines)
		fmt.Println(ans)
		ans = QuerySerially(word, engines)
		fmt.Println(ans)
	*/
}
