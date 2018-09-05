package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"zk"
)

func main() {

	var hosts = []string{"192.168.30.19:2181"} //zk服务器地址
	var zk_server_root_path = "/server"        //zk根目录
	var server_addr = "192.168.19.30"          //开启服务地址
	var server_port = ":65530"                 //服务占用端口

	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建根目录 权限为所有权限
	conn.Create(zk_server_root_path, nil, 0, zk.WorldACL(zk.PermAll))

	//注册服务
	var path = zk_server_root_path + "/" + server_addr + server_port
	const flags int32 = 1
	//0:永久，除非手动删除
	//zk.FlagEphemeral = 1:短暂，session断开则改节点也被删除
	//zk.FlagSequence  = 2:会自动在节点后面添加序号
	//3:Ephemeral和Sequence，即，短暂且自动添加序号
	//只读权限
	var acls = zk.WorldACL(zk.PermRead)
	p, err_create := conn.Create(path, nil, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Println("create:", p)

	// 启动web服务，监听9090端口
	http.HandleFunc("/", index)
	listen_err := http.ListenAndServe(server_port, nil)
	if listen_err != nil {
		log.Fatal("ListenAndServe: ", listen_err)
	}

	defer conn.Close()
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出

	fmt.Fprintf(w, "server time "+time.Now().Format("2006-01-02 15:04:05"))
}
