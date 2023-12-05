package main

import (
	"math/rand"
	"strconv"
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
	Data []int64 `json:"data"`
}

func randomInt(c *gin.Context) {
	maxS := c.Query("max")
	minS := c.Query("min")
	numS := c.Query("num")

	max, _ := strconv.ParseInt(maxS, 10, 32)
	min, _ := strconv.ParseInt(minS, 10, 32)
	num, _ := strconv.ParseInt(numS, 10, 32)

	var data []int64
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(num); i++ {
		randomNum := min + rand.Int63n(max-min)
		data = append(data, randomNum)
	}

	resp := &RandomResp{
		Data: data,
	}

	c.JSON(200, resp)
}
