package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建一个默认的路由引擎
var router = gin.Default()

func Register() {
	// 定义一个GET请求的路由处理函数
	router.GET("/", randomInt)

}

type RandomResp struct {
	IntData    int64  `json:"intdata"`
	StringData string `json:"stringdata"`
	Type       int    `json:"type"`
}

func randomInt(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	resp := &RandomResp{}

	typ := rand.Intn(2)
	if typ >= 1 {
		randomNum := 0 + rand.Int63n(100)
		resp.IntData = randomNum
		resp.Type = 2
	} else {
		randomStr := "hello"
		resp.StringData = randomStr
		resp.Type = 1
	}

	c.JSON(200, resp)
}
