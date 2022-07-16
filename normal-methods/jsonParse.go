package main

import (
	"encoding/json"
	"fmt"
)

/**
需要注意的点：
- 只有公有对象或者属性才会可以转json
- 字段tag是”-“, 则该字段是不会输出到json 的
- 另一个json包 github.com/bitly/go-simplejson
*/

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

type Person struct {
	ID     int      `json:"-"` // 不会输出
	Name   string   `json:"name"`
	Salary float64  `json:"salary,string"`   // 强转成string
	Hobby  []string `json:"hobby,omitempty"` // 如果为空，字段不会在json中
}

func main() {
	// parse json to obj
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println("obj => ", s)

	// parse obj to json
	s = Serverslice{}
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("json", string(b))

	// transfer obj to json and use special Name
	teacher := Person{
		ID:     1,
		Name:   "teacher",
		Salary: 10111.223,
		//Hobby:  []string{"running", "swimming"},
		Hobby: nil,
	}
	fmt.Println(teacher)
	res, err := json.Marshal(teacher)
	fmt.Println(string(res))
}
