package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	// 这里是路径路由，即接受/api/websocket路径的websocket请求
	r.GET("/api/websocket", func(c *gin.Context) {
		proxyWsRequest(c.Writer, c.Request)
	})

	// 这里是处理其他请求的路由
	r.NoRoute(func(c *gin.Context) {
		// 创建一个反向代理
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "https", // 代理的源站地址，需要根据自己的实际情况修改
			Host:   "example.com",
		})

		// 修改Header的Host和Referer
		c.Request.Header.Set("Host", "example.com")
		c.Request.Header.Set("Referer", "https://example.com")

		// 请求代理
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// 这里是启动服务
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func proxyWsRequest(w http.ResponseWriter, r *http.Request) {
	// 将wss升级为websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级websocket失败：", err)
		return
	}
	defer conn.Close()

	// 获取wss服务器的连接地址
	wssAddr := "wss://example.com/ws" // 根据实际情况修改
	wssUrl, err := url.Parse(wssAddr)
	if err != nil {
		fmt.Println("解析wss地址失败：", err)
		return
	}
}
