package main

import (
	"common/model"
	"common/tool"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	// 必须要在数据库里面存在
	a := tool.JwtUtil.CreateTokenEasy(3)
	b := tool.JwtUtil.CreateTokenEasy(12313)
	c := tool.JwtUtil.CreateTokenEasy(12314)
	m := map[int]string{
		0: a,
		1: b,
		2: c,
	}
	var group sync.WaitGroup
	for i := 0; i < 20; i++ {
		//group.Add(1)
		//go func(i int) {
		//	defer group.Done()
		req, _ := http.NewRequest("GET", "http://localhost:9999/admin/getUserInfo", nil)
		req.Header.Add("Authorization", m[i%3])
		// 发送请求
		client := &http.Client{}
		resp, _ := client.Do(req)
		// 读取响应
		data, _ := io.ReadAll(resp.Body)
		var result model.Result
		_ = json.Unmarshal(data, &result)
		// 这里看到应该要与请求token中的id一致，否则会出现问题
		fmt.Println(result.Data)
		//}(i)
	}
	group.Wait()

}
