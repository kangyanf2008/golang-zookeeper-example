package main

import (
	"fmt"
	"time"
	"zk"
)

func main() {
	var hosts = []string{"192.168.30.19:2181"} //server端host
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	children, stat, ch, err := conn.ChildrenW("/server")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)

	var server_list []server_info

	if len(children) > 0 {
		serverNum := len(children)
		var server_list2 [7]server_info
		for i := 0; i < len(children); i++ {
			fmt.Println(children[i])
			server_list2[i] = server_info{children[i], 0}
		}
		server_list := &server_list2
	}
	//v_server_info := server_info{children[1], 0}
	//fmt.Print(v_server_info.serverAddress)
	//fmt.Println(server_info{children[0], 0})
	fmt.Println(server_list[0])
	//server_list[0] = server_info{children[0], 0}

	if ch != nil {
		e := <-ch
		fmt.Printf("%+v\n", e)
	}

	defer conn.Close()
}
