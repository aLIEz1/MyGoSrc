package main

func main()  {
	//defer延迟调用 mian函数结束前调用
	//defer先写的后调用(先进后出)哪怕某一个函数或者延迟调用发生错误，这些调用依旧会被执行
	defer println("bbb")
	defer println("aaa")
	defer println("ccc")

}