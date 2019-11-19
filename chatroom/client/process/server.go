package process

import (
	"fmt"
	"go_project/chatroom/server/utils"
	"net"
	"os"
)

// 显示登陆成功后的界面.....................
func ShowMenu() {
	fmt.Println("-------------恭喜登陆成功----------------------")
	fmt.Println("-------------1、显示在线用户列表----------------")
	fmt.Println("-------------2、发送消息----------------------")
	fmt.Println("-------------3、信息列表-----------------------")
	fmt.Println("-------------4、退出系统-----------------------")
	fmt.Println("清选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
		case 1:
			fmt.Println("显示在线用户列表...")
		case 2:
			fmt.Println("ddasdsadqw")
		case 4:
			fmt.Println("你选择了退出系统.........")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误！！！！")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {

	// 创建一个transfer实例不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端持续读取服务端的链接出错了", err)
			return
		}
		//如果读到消息
		fmt.Printf("mes=%v", mes)
	}
}