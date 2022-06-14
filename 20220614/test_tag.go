// json中使用-编解码
package main

import (
	"encoding/json"
	"fmt"
)

// 定义多类型结构体，标签``
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Action []string `json:"actors"`
}

func main() {
	// 初始化结果体
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}
	// struct转json数据--marshal--err错误提示
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr=%s\n", jsonStr)
	fmt.Println("------------------")
	// 将上面的json数据赋值给myMovie
	myMovie := Movie{}
	// json数据转结构体,指针传递
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	fmt.Printf("%v\n", myMovie)

}
