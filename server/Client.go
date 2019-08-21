package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("nert.Dial err = ",err)
		return
	}
	//main调用完毕，关闭连接
	defer conn.Close()

	go func() {

		//从键盘输入内容，给服务器发送内容
		for  {
			str:=make([]byte,1024)
			n,err3:=os.Stdin.Read(str)   //从键盘读取内容，放在str
			if err3!=nil{
				fmt.Println("os.Stadin.err",err3)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])


		}

	}()
	//接收服务器的回复信息
	//切片缓冲
	buf:=make([]byte,1024)
	for{
		n,err1:=conn.Read(buf)
		if err!=nil {
			fmt.Println("conn.Read err1 = ",err1)
			return

		}
		fmt.Println(string(buf[:n]))  //打印接收到的内容,转换为字符串打印

	}
}