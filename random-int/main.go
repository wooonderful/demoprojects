package main

func main() {
	Register()
	// 启动HTTP服务，监听在本地的8080端口
	router.Run(":38080")
}
