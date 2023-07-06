package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建一个默认的路由引擎
var router = gin.Default()

func Register() {
	// 定义一个GET请求的路由处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.GET("/api/v2/version", version)
}

var count = 0

type NeuronVersionResp struct {
	BuildDate string `json:"build_date"`
	Revision  string `json:"revision"`
	Version   string `json:"version"`
}

func version(c *gin.Context) {
	count++
	ret := &NeuronVersionResp{
		BuildDate: "2023-7-5",
		Revision:  "2023-7-5",
		Version:   "2.3.6",
	}

	if count%10 == 0 {
		durationCount(func() { time.Sleep(time.Duration(10+rand.Intn(30)) * time.Second) })
		c.JSON(200, ret)
	} else {
		c.JSON(200, ret)
	}
}

func durationCount(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	cost := end.Unix() - start.Unix()
	fmt.Printf("execute function cost: %d s\n", cost)
}
